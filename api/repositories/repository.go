package repositories

import (
	"clean-architecture/db/models"
	"context"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	ListOrders(ctx context.Context) ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) ListOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
