package service

import (
	"go-api/model"
	"go-api/repository"
	"fmt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id int, user model.User) error {
	return s.Repo.UpdateUser(id, user)
}

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error finding user by email: %v", err)
	}
	return user, nil
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("error finding user by id: %v", err)
	}
	return user, nil
}