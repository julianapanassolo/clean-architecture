package order

import "time"

// Order representa um pedido
type Order struct {
	ID          int32     `db:"id"`
	CustomerID  int32     `db:"customer_id"`
	TotalAmount float64   `db:"total_amount"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
