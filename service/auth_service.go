package service

import (
	"go-api/model"
	"go-api/utils"
	"fmt"
	"log" 
)

type AuthService struct {
	UserService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{UserService: userService}
}

func (s *AuthService) RegisterUser(user model.User) error {
	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("email and password are required")
	}
	return s.UserService.CreateUser(user)
}

func (s *AuthService) ValidateUser(email, password string) (bool, error) {
	if email == "" || password == "" {
		return false, fmt.Errorf("email and password cannot be empty")
	}
	log.Println(email) 
	user, err := s.UserService.Repo.FindByEmail(email)
	log.Println(user, err) 
	if err != nil {
		return false, fmt.Errorf("user not found")
	}
	if user.Password != password {
		return false, fmt.Errorf("invalid password")
	}
	return true, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	isValid, err := s.ValidateUser(email, password)
	if err != nil || !isValid {
		return "", fmt.Errorf("invalid email or password")
	}
	token, err := utils.GenerateJWT(email)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	return token, nil
}

