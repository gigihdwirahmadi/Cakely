package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (c *AuthController) RegisterUser(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	if err := c.Service.RegisterUser(user); err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to register user")
	}
	return utils.Respond(ctx, http.StatusCreated, nil, "User registered successfully")
}

func (c *AuthController) Login(ctx echo.Context) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&credentials); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid input")
	}

	token, err := c.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		return utils.Respond(ctx, http.StatusUnauthorized, nil, "Invalid email or password")
	}

	return utils.Respond(ctx, http.StatusOK, map[string]string{"token": token}, "Login successful")
}

func (c *AuthController) Logout(ctx echo.Context) error {
	user := ctx.Get("user")
	if user == nil {
		return utils.Respond(ctx, http.StatusUnauthorized, nil, "Unauthorized")
	}
	//untuk logout logic nya lewat FE dengan tidak parsing token
	return utils.Respond(ctx, http.StatusOK, nil, "Logout successful")
}
