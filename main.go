package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/router"
	"go-api/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize database connection
	config.InitDB()
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())   // Logging requests
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB stack trace
	}))

	// Initialize repositories
	userRepo := repository.NewUserRepository(config.DB)
	cakeRepo := repository.NewCakeRepository(config.DB)
	orderRepo := repository.NewOrderRepository(config.DB)

	// Initialize services
	userService := service.NewUserService(userRepo)
	cakeService := service.NewCakeService(cakeRepo)
	orderService := service.NewOrderService(orderRepo, cakeService, userService)
	authService := service.NewAuthService(userService)

	// Initialize controllers
	authController := controller.NewAuthController(authService)
	cakeController := controller.NewCakeController(cakeService)
	orderController := controller.NewOrderController(orderService)

	// Initialize routers
	router.InitRouter(e, authController, cakeController, orderController)

	// Start the server
	e.Logger.Fatal(e.Start(":8089"))
}
