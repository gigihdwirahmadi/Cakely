package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Service *service.OrderService
}

func NewOrderController(service *service.OrderService) *OrderController {
	return &OrderController{Service: service}
}

func (c *OrderController) CreateOrder(ctx echo.Context) error {
	var order model.Order
	if err := ctx.Bind(&order); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	if err := c.Service.CreateOrder(order); err != nil {
		log.Println(err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to create order")
	}
	return utils.Respond(ctx, http.StatusCreated, nil, "Order created successfully")
}

func (c *OrderController) GetOrderByID(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid order ID")
	}
	order, err := c.Service.GetOrderByID(id)
	if err != nil {
		return utils.Respond(ctx, http.StatusNotFound, nil, "Order not found")
	}
	response, err := c.Service.MapOrderToResponse(order)
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to map order response")
	}
	return utils.Respond(ctx, http.StatusOK, response, "Order retrieved successfully")
}

func (c *OrderController) GetAllOrders(ctx echo.Context) error {
	orders, err := c.Service.GetAllOrders()
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve orders")
	}
	responses, err := c.Service.MapOrdersToResponses(orders)
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to map orders response")
	}
	return utils.Respond(ctx, http.StatusOK, responses, "Orders retrieved successfully")
}

func (c *OrderController) DeleteOrder(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid order ID")
	}
	if err := c.Service.DeleteOrder(id); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to delete order")
	}
	return utils.Respond(ctx, http.StatusOK, nil, "Order deleted successfully")
}