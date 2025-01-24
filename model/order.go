package model

import "time"

type Order struct {
	ID          int       `db:"id" json:"id"`
	CakeID      int       `db:"cake_id" json:"cake_id"`
	UserID      int       `db:"user_id" json:"user_id"`
	Quantity    int       `db:"quantity" json:"quantity"`
	Total       int    `db:"total" json:"total"`
	TrxDateTime time.Time `db:"trx_datetime" json:"trx_datetime"`
}

