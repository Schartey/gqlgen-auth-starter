package main

import (
	"github.com/schartey/gqlgen-auth-starter/gqlgen/resolvers"
	"github.com/schartey/gqlgen-auth-starter/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/schartey/gqlgen-auth-starter/gqlgen"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	users := map[string]*model.User{
		"1": {
			ID:       "1",
			Username: "Joe",
			Person: model.Person{
				ID:        "1",
				Firstname: "John",
				Lastname:  "Doe",
				Email:     "john.doe@mail.com",
				Phone:     "+1234567890",
				Birthdate: time.Now(),
			},
		},
		"2": {
			ID:       "2",
			Username: "Jane",
			Person: model.Person{
				ID:        "2",
				Firstname: "Jane",
				Lastname:  "Doe",
				Email:     "jane.doe@mail.com",
				Phone:     "+1345678901",
				Birthdate: time.Now(),
			},
		},
	}

	resolver := resolvers.NewRootResolver(users)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
