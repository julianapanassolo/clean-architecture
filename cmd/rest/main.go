package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"clean-architecture/pkg/database"
	"clean-architecture/pkg/service"

	"github.com/gorilla/mux"
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

	// Define o roteador
	router := mux.NewRouter()

	// Define o endpoint REST para listar os pedidos
	router.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		// Chama o serviço para listar os pedidos
		orders, err := orderService.ListOrders(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("erro ao listar os pedidos: %v", err), http.StatusInternalServerError)
			return
		}

		// Serializa os pedidos para JSON
		json.NewEncoder(w).Encode(orders)
	})

	// Inicia o servidor HTTP na porta 9030
	fmt.Printf("Servidor REST iniciado na porta 9030\n")
	log.Fatal(http.ListenAndServe(":9030", router))
}
