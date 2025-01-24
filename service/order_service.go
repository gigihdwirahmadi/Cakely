package service

import (
	"go-api/model"
	"go-api/response"
	"go-api/repository"
	"fmt"
)

type OrderService struct {
	Repo         *repository.OrderRepository
	CakeService  *CakeService
	UserService  *UserService
}

func NewOrderService(repo *repository.OrderRepository, cakeService *CakeService, userService *UserService) *OrderService {
	return &OrderService{Repo: repo, CakeService: cakeService, UserService: userService}
}

func (s *OrderService) CreateOrder(order model.Order) error {
	return s.Repo.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(id int) (*model.Order, error) {
	return s.Repo.GetOrderByID(id)
}

func (s *OrderService) GetAllOrders() ([]model.Order, error) {
	return s.Repo.GetAllOrders()
}

func (s *OrderService) DeleteOrder(id int) error {
	return s.Repo.DeleteOrder(id)
}

func (s *OrderService) MapOrderToResponse(order *model.Order) (*response.OrderResponse, error) {
	cake, err := s.CakeService.GetCakeByID(order.CakeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cake name: %v", err)
	}
	user, err := s.UserService.GetUserByID(order.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user name: %v", err)
	}
	return &response.OrderResponse{
		ID:          order.ID,
		CakeName:    cake.Name,
		UserName:    user.Name,
		Quantity:    order.Quantity,
		Total:       order.Total,
		TrxDateTime: order.TrxDateTime,
	}, nil
}

func (s *OrderService) MapOrdersToResponses(orders []model.Order) ([]response.OrderResponse, error) {
	responses := make([]response.OrderResponse, len(orders))
	for i, order := range orders {
		response, err := s.MapOrderToResponse(&order)
		if err != nil {
			return nil, err
		}
		responses[i] = *response
	}
	return responses, nil
}
