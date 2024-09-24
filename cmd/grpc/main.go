package main

import (
	"fmt"
	"log"
	"net"

	"clean-architecture/grpc/pb"
	"clean-architecture/pkg/database"
	"clean-architecture/pkg/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Define a URL de conexão com o banco de dados MySQL
	dbURL := "mysql://root:social@localhost:9030/social?charset=utf8mb4&parseTime=true&loc=Local"

	// Cria uma nova instância de Database
	db, err := database.NewDatabase(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cria uma nova instância de OrderService
	orderService := service.NewOrderService(db)

	// Cria um novo servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Cria um novo servidor gRPC
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &pb.OrderServiceServer{
		OrderService: orderService,
	})

	// Registra o serviço de reflexão para o gRPC
	reflection.Register(s)

	fmt.Printf("Servidor gRPC iniciado na porta 50051\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
