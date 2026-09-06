package domain

import "time"

type Order struct {
	ID           int64     `json:"id"`
	OrderNumber  string    `json:"order_number"`
	CustomerName string    `json:"customer_name"`
	Amount       float64   `json:"amount"`
	SyncStatus   string    `json:"sync_status"`
	Attempts     int       `json:"attempts"`
	LastError    *string   `json:"last_error"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
