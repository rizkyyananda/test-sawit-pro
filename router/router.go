package router

import (
	"github.com/labstack/echo/v4"
	"test_sawit_pro/controller"
)

type Handlers struct {
	EstateController controller.EstateController
}

func RegisterRoutes(e *echo.Echo, h *Handlers) {
	api := e.Group("/api")

	RegisterEstateRoutes(api, h.EstateController)
}
