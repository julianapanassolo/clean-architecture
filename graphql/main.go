package graphql

import (
	"clean-architecture/db"
	"clean-architecture/graphql/resolvers"

	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/codegen/testserver/nullabledirectives/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func StartGraphQLServer() {
	db := db.GetDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("GraphQL server listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
