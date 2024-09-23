package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderServer struct {
	pb.UnimplementedOrderServer
	orderDB *db.db
}

func NewOrderServer(orderDB *db.db) *OrderServer {
	return &OrderServer{
		orderDB: orderDB,
	}
}

func (os *OrderServer) ListOrders(ctx context.Context, req *pb.Empty) (*pb.ListOrdersResponse, error) {
	orders, err := os.orderDB.ListOrders()
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to retrieve orders from database")
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:          order.ID,
			CustomerId:  order.CustomerID,
			Products:    order.Products,
			TotalAmount: order.TotalAmount,
		})
	}

	return &pb.ListOrdersResponse{
		Orders: pbOrders,
	}, nil
}
