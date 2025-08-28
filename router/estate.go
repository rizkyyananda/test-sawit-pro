package router

import (
	"github.com/labstack/echo/v4"
	"test_sawit_pro/controller"
)

func RegisterEstateRoutes(g *echo.Group, h controller.EstateController) {
	estate := g.Group("/estate")

	estate.POST("", h.CreateEstate)
	estate.POST("/:id/tree", h.AddTree)
	estate.GET("/:id/stats", h.GetTreeStats)
	estate.GET("/:id/drone-plan", h.GetDronePlan)
	estate.GET("/max/:id/drone-plan", h.MaxDistance)
}
