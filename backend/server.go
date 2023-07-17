package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
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

	//SQLBoilerのデバッグ出力を有効化
	boil.DebugMode = true

	service := services.New(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Srv:     service,
			Loaders: graph.NewLoaders(service),
		},
		Directives: graph.Directive,
		Complexity: graph.ComplexityConfig(),
	}))
	//複雑度10以上のクエリは発行できないようにする
	srv.Use(extension.FixedComplexityLimit(10))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	// SQLBoilerによって発行されるSQLクエリをログ出力させるデバッグオプション
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
