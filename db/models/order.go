package models

type Order struct {
	ID          int      `json:"id"`
	CustomerID  int      `json:"customer_id"`
	Products    []string `json:"products"`
	TotalAmount float64  `json:"total_amount"`
}
