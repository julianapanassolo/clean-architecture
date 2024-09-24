package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"clean-architecture/pkg/database"
	"clean-architecture/pkg/service"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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

	// Define o esquema GraphQL
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"customerId": &graphql.Field{
				Type: graphql.Int,
			},
			"totalAmount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"listOrders": &graphql.Field{
				Type: graphql.NewList(orderType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					ctx := context.Background()
					orders, err := orderService.ListOrders(ctx)
					if err != nil {
						return nil, err
					}
					return orders, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Cria um novo handler GraphQL
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Inicia o servidor GraphQL na porta 9031
	fmt.Printf("Servidor GraphQL iniciado na porta 9031\n")
	log.Fatal(http.ListenAndServe(":9031", h))
}
