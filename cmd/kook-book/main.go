package main

import (
	"log"
	"net/http"
	"os"

	"vanjiii/kook-book-server/pkg/graphql"
	"vanjiii/kook-book-server/pkg/services"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	env, err := services.NewEnv()
	if err != nil {
		log.Fatalf("error occurs on creating environment: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	res := graphql.NewResolver(env)

	// TODO: remove this
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write(services.Page)
	})

	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: res})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
