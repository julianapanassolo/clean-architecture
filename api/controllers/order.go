package controllers

import (
	"clean-architecture/db/models"

	"context"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type OrderController struct {
	grpcClient pb.OrderServiceClient
}

func NewOrderController() *OrderController {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &OrderController{
		grpcClient: pb.NewOrderServiceClient(conn),
	}
}

func (oc *OrderController) ListOrders(c *gin.Context) {
	ctx := context.Background()
	res, err := oc.grpcClient.ListOrders(ctx, &pb.Empty{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var orders []models.Order
	for _, order := range res.Orders {
		orders = append(orders, models.Order{
			ID:          int(order.Id),
			CustomerID:  int(order.CustomerId),
			Products:    order.Products,
			TotalAmount: float64(order.TotalAmount),
		})
	}

	c.JSON(200, gin.H{"orders": orders})
}
