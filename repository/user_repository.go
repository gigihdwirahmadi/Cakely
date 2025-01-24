package repository

import (
	"go-api/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user model.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) UpdateUser(id int, user model.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, id)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	query := "SELECT id, name, email, password FROM users WHERE email = $1 LIMIT 1"
	err := r.DB.Get(&user, query, email)
	return &user, err
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	query := "SELECT id, name, email, password FROM users WHERE id = $1"
	err := r.DB.Get(&user, query, id)
	return &user, err
}
