package grpc

import (
	"clean-architecture/db"
	"clean-architecture/grpc/pb"
	"clean-architecture/grpc/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServer(s, server.NewOrderServer(db.GetDB()))
	fmt.Println("GRPC server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
