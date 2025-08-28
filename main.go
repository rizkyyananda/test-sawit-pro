package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"test_sawit_pro/config"
	"test_sawit_pro/depedency-injection"
	"test_sawit_pro/router"
	"time"
)

func main() {
	isDebug := true
	envFlag := flag.String("env", "", "App environment: local|staging|production")
	flag.Parse()
	env := *envFlag
	if env == "" {
		env = os.Getenv("APP_ENV")
	}
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatal("load config:", err)
	}

	container, err := depedency_injection.Init(cfg)
	if err != nil {
		log.Fatal("dependency init:", err)
	}

	e := echo.New()
	e.Debug = isDebug
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		status := map[string]interface{}{
			"status":  "UP",
			"service": "Sawit pro service is running",
			"time":    time.Now().Format(time.RFC3339),
		}
		return c.JSON(http.StatusOK, status)
	})

	// ========== Tambahan untuk Dokumentasi ==========
	e.File("/openapi.yaml", "api.yaml")

	// Swagger UI
	e.GET("/docs", func(c echo.Context) error {
		html := `<!DOCTYPE html>
		<html>
		<head>
		  <meta charset="utf-8"/>
		  <title>SwaggerUI</title>
		  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css">
		</head>
		<body>
		<div id="swagger-ui"></div>
		<script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
		<script>
		  window.ui = SwaggerUIBundle({
			url: '/openapi.yaml',
			dom_id: '#swagger-ui'
		  });
		</script>
		</body>
		</html>`
		return c.HTML(http.StatusOK, html)
	})
	// ==================================================

	// Daftarkan routes aplikasi
	router.RegisterRoutes(e, &router.Handlers{
		EstateController: *container.EstateController,
	})

	port := cfg.Server.Port
	addr := ":" + strconv.Itoa(port)
	log.Println("Server running on", addr)

	// Start server
	if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
		log.Fatal("server error:", err)
	}
}
