package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"task-managament/server/api/graph/generated"
	"task-managament/server/api/graph/resolvers"
	"task-managament/server/commons"
	"task-managament/server/repositories/task_management"
	"task-managament/server/services/task"
)

const defaultPort = "8080"

func main() {

	repo, err := taskmanagement.GetRepository(commons.RepositoryTypeDB)
	taskManagementService := task.New(repo)
	resolver := resolvers.Resolver{
		TasksManagementService: taskManagementService,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))
	http.Handle(commons.EndpointPath, srv)

	ctx := context.Background()
	if err != nil {
		log.Fatal(ctx, err, "Unable to get repository")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
