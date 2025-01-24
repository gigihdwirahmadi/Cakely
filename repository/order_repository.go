package repository

import (
	"go-api/model"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	DB *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(order model.Order) error {
	query := "INSERT INTO orders (cake_id, quantity, total, user_id, trx_datetime) VALUES ($1, $2, $3, $4, NOW())"
	_, err := r.DB.Exec(query, order.CakeID, order.Quantity, order.Total, order.UserID)
	return err
}

func (r *OrderRepository) GetOrderByID(id int) (*model.Order, error) {
	var order model.Order
	query := "SELECT id, cake_id, quantity, total, user_id, trx_datetime FROM orders WHERE id = $1"
	err := r.DB.Get(&order, query, id)
	return &order, err
}

func (r *OrderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	query := "SELECT id, cake_id, quantity, total, user_id, trx_datetime FROM orders"
	err := r.DB.Select(&orders, query)
	return orders, err
}

func (r *OrderRepository) DeleteOrder(id int) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
