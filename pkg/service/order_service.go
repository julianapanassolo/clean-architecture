package service

import (
	"clean-architecture/pkg/database"
	"clean-architecture/pkg/order"
	"context"
	"fmt"
)

// OrderService é o serviço de pedidos
type OrderService struct {
	db *database.Database
}

// NewOrderService cria uma nova instância de OrderService
func NewOrderService(db *database.Database) *OrderService {
	return &OrderService{db: db}
}

// ListOrders lista todos os pedidos
func (s *OrderService) ListOrders(ctx context.Context) ([]*order.Order, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, fmt.Errorf("erro ao listar os pedidos: %w", err)
	}
	defer rows.Close()

	var orders []*order.Order
	for rows.Next() {
		var o order.Order
		if err := rows.Scan(&o.ID, &o.CustomerID, &o.TotalAmount, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, fmt.Errorf("erro ao ler os pedidos: %w", err)
		}
		orders = append(orders, &o)
	}

	return orders, nil
}
