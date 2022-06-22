package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"task-managament/graph"
	"task-managament/graph/generated"
	"task-managament/server/commons"
	"task-managament/server/repositories"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	_, err := repositories.GetRepository(commons.RepositoryTypeDB)
	ctx := context.Background()
	if err != nil {
		log.Fatal(ctx, err, "Unable to get repository")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
