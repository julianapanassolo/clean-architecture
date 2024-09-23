package resolvers

import (
	"context"
)

func (r *Resolver) Orders(ctx context.Context) ([]*generated.Order, error) {
	orders, err := r.DB.ListOrders()
	if err != nil {
		return nil, err
	}

	var gqlOrders []*generated.Order
	for _, order := range orders {
		gqlOrders = append(gqlOrders, &generated.Order{
			ID:          int(order.ID),
			CustomerID:  int(order.CustomerID),
			Products:    order.Products,
			TotalAmount: order.TotalAmount,
		})
	}

	return gqlOrders, nil
}
