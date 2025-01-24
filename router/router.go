package router

import (
	"go-api/controller"
	"go-api/utils"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, authController *controller.AuthController,  cakeController *controller.CakeController, orderController *controller.OrderController) {
	e.POST("/auth/register", authController.RegisterUser)
	e.POST("/auth/login", authController.Login)
	e.POST("/auth/logout", authController.Logout)

	// Protected routes
	protected := e.Group("")
	protected.Use(utils.JWTMiddleware())

	protected.POST("/cakes", cakeController.CreateCake)
	protected.PUT("/cakes/:id", cakeController.UpdateCake)
	protected.GET("/cakes/:id", cakeController.GetCakeByID)
	protected.GET("/cakes", cakeController.GetAllCakes)
	protected.DELETE("/cakes/:id", cakeController.DeleteCake)

	protected.POST("/orders", orderController.CreateOrder)
	protected.GET("/orders/:id", orderController.GetOrderByID)
	protected.GET("/orders", orderController.GetAllOrders)
	protected.DELETE("/orders/:id", orderController.DeleteOrder)
}