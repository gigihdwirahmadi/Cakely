package repository

import (
	"github.com/jmoiron/sqlx"
	"go-api/model"
)

type CakeRepository struct {
	DB *sqlx.DB
}

func NewCakeRepository(db *sqlx.DB) *CakeRepository {
	return &CakeRepository{DB: db}
}

func (r *CakeRepository) CreateCake(cake model.Cake) error {
	query := "INSERT INTO cakes (name, price, description) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, cake.Name, cake.Price, cake.Description)
	return err
}

func (r *CakeRepository) UpdateCake(id int, cake model.Cake) error {
	query := "UPDATE cakes SET name = $1, price = $2, description = $3 WHERE id = $4"
	_, err := r.DB.Exec(query, cake.Name, cake.Price, cake.Description, id)
	return err
}

func (r *CakeRepository) GetCakeByID(id int) (*model.Cake, error) {
	var cake model.Cake
	query := "SELECT id, name, price, description FROM cakes WHERE id = $1"
	err := r.DB.Get(&cake, query, id)
	return &cake, err
}

func (r *CakeRepository) GetAllCakes() ([]model.Cake, error) {
	var cakes []model.Cake
	query := "SELECT id, name, price, description FROM cakes"
	err := r.DB.Select(&cakes, query)
	return cakes, err
}

func (r *CakeRepository) DeleteCake(id int) error {
	query := "DELETE FROM cakes WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}