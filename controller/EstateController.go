package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test_sawit_pro/dto/request"
	"test_sawit_pro/entity"
	"test_sawit_pro/pkg/helper"
	"test_sawit_pro/service"
)

var response = helper.Response{}

type EstateController interface {
	CreateEstate(ctx echo.Context) error
	AddTree(ctx echo.Context) error
	GetTreeStats(ctx echo.Context) error
	GetDronePlan(ctx echo.Context) error
	MaxDistance(ctx echo.Context) error
}

type estateController struct {
	estateService service.EstateService
	treeService   service.TreeService
	droneService  service.DroneService
}

func NewEstateController(es service.EstateService, ts service.TreeService, ds service.DroneService) EstateController {
	return &estateController{
		estateService: es,
		treeService:   ts,
		droneService:  ds,
	}
}

func (e *estateController) MaxDistance(ctx echo.Context) error {
	var response helper.Response
	estateID := ctx.Param("id")
	maxDistance, _ := strconv.Atoi(ctx.QueryParam("max_distance"))

	res, err, statusCode := e.droneService.MaxDistance(ctx.Request().Context(), estateID, &maxDistance)
	if err != nil {
		response = helper.ResponseError(err.Error(), statusCode)
		return ctx.JSON(statusCode, response)
	}
	response = helper.ResponseSuccess(res)
	return ctx.JSON(statusCode, response)
}

func (e *estateController) CreateEstate(ctx echo.Context) error {
	var req request.CreateEstateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseError("invalid input", http.StatusBadRequest))
	}

	res, err, statusCode := e.estateService.CreateEstate(ctx.Request().Context(), &entity.Estate{
		Width:  req.Width,
		Length: req.Length,
	})

	if err != nil {
		return ctx.JSON(statusCode, helper.ResponseError(err.Error(), statusCode))
	}
	return ctx.JSON(statusCode, helper.ResponseSuccess(res))
}

func (e *estateController) AddTree(ctx echo.Context) error {
	estateID := ctx.Param("id")

	var req request.AddTreeRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ResponseError("invalid input", http.StatusBadRequest))
	}

	res, err, statusCode := e.treeService.AddTree(ctx.Request().Context(), &entity.Tree{
		Height:   req.Height,
		Y:        req.Y,
		X:        req.X,
		EstateID: estateID,
	})

	if err != nil {
		return ctx.JSON(statusCode, helper.ResponseError(err.Error(), statusCode))
	}
	return ctx.JSON(statusCode, helper.ResponseSuccess(res))
}

func (e *estateController) GetTreeStats(ctx echo.Context) error {
	estateID := ctx.Param("id")

	res, err, statusCode := e.treeService.GetTreeStats(ctx.Request().Context(), estateID)
	if err != nil {
		return ctx.JSON(statusCode, helper.ResponseError(err.Error(), statusCode))
	}
	return ctx.JSON(statusCode, helper.ResponseSuccess(res))
}

func (e *estateController) GetDronePlan(ctx echo.Context) error {
	estateID := ctx.Param("id")

	res, err, statusCode := e.droneService.CalculateDroneDistance(ctx.Request().Context(), estateID)
	if err != nil {
		return ctx.JSON(statusCode, helper.ResponseError(err.Error(), statusCode))
	}
	return ctx.JSON(statusCode, helper.ResponseSuccess(res))
}
