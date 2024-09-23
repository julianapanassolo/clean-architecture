package main

import (
	"clean-architecture/api"
	"clean-architecture/graphql"
	"clean-architecture/grpc"
	"log"
	"net/http"
)

func main() {
	// Iniciar o servidor GRPC
	go grpc.StartGRPCServer()

	// Iniciar o servidor GraphQL
	go graphql.StartGraphQLServer()

	// Iniciar o servidor REST API
	apiServer := &http.Server{
		Addr:    ":9030",
		Handler: api.NewRouter(),
	}

	log.Fatal(apiServer.ListenAndServe())
}
