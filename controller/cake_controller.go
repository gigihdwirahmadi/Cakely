package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"net/http"
	"strconv"
	"log"

	"github.com/labstack/echo/v4"
)

type CakeController struct {
	Service *service.CakeService
}

func NewCakeController(service *service.CakeService) *CakeController {
	return &CakeController{Service: service}
}

func (c *CakeController) CreateCake(ctx echo.Context) error {
	var cake model.Cake
	if err := ctx.Bind(&cake); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	if err := c.Service.CreateCake(cake); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to create cake")
	}
	return utils.Respond(ctx, http.StatusCreated, nil, "Cake created successfully")
}

func (c *CakeController) UpdateCake(ctx echo.Context) error {
	var cake model.Cake
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid cake ID")
	}
	if err := ctx.Bind(&cake); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	if err := c.Service.UpdateCake(id, cake); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to update cake")
	}
	return utils.Respond(ctx, http.StatusOK, nil, "Cake updated successfully")
}

func (c *CakeController) GetCakeByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid cake ID")
	}
	cake, err := c.Service.GetCakeByID(id)
	log.Println(cake, err) 
	if err != nil {
		return utils.Respond(ctx, http.StatusNotFound, nil, "Cake not found")
	}
	return utils.Respond(ctx, http.StatusOK, cake, "Cake retrieved successfully")
}

func (c *CakeController) GetAllCakes(ctx echo.Context) error {
	cakes, err := c.Service.GetAllCakes()
	log.Println(cakes, err) 
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve cakes")
	}
	return utils.Respond(ctx, http.StatusOK, cakes, "Cakes retrieved successfully")
}

func (c *CakeController) DeleteCake(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid cake ID")
	}
	if err := c.Service.DeleteCake(id); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to delete cake")
	}
	return utils.Respond(ctx, http.StatusOK, nil, "Cake deleted successfully")
}
