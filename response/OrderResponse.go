package response

import "time"

type OrderResponse struct {
	ID          int       `json:"id"`
	CakeName    string    `json:"cake_name"`
	UserName    string    `json:"user_name"`
	Quantity    int       `json:"quantity"`
	Total       int    `json:"total"`
	TrxDateTime time.Time `json:"trx_datetime"`
}