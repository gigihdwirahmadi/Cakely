package service

import (
	"go-api/model"
	"go-api/repository"
)

type CakeService struct {
	Repo *repository.CakeRepository
}

func NewCakeService(repo *repository.CakeRepository) *CakeService {
	return &CakeService{Repo: repo}
}

func (s *CakeService) CreateCake(cake model.Cake) error {
	return s.Repo.CreateCake(cake)
}

func (s *CakeService) UpdateCake(id int, cake model.Cake) error {
	return s.Repo.UpdateCake(id, cake)
}

func (s *CakeService) GetCakeByID(id int) (*model.Cake, error) {
	return s.Repo.GetCakeByID(id)
}

func (s *CakeService) GetAllCakes() ([]model.Cake, error) {
	return s.Repo.GetAllCakes()
}

func (s *CakeService) DeleteCake(id int) error {
	return s.Repo.DeleteCake(id)
}