package model

type Cake struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Price int `db:"price" json:"price"`
	Description   string    `db:"description" json:"description"`
}