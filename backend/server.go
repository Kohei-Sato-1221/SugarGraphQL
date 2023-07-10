package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/graph"
	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/graph/services"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("mysql", "root:pass@(localhost:3306)/graphqldb?parseTime=true")
	if err != nil {
		panic("failed to get db connection")
	}
	defer db.Close()

	service := services.New(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Srv: service,
		},
		Directives: graph.Directive,
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
