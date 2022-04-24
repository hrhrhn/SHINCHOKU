package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"dev-gorm-gql-v3/backend"

	"github.com/99designs/gqlgen/handler"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

const defaultPort = "5050"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		log.Fatalln("接続失敗", err)
	} else {
		fmt.Println("--init--")
	}
	defer db.Close()
	db.LogMode(true)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(backend.NewExecutableSchema(backend.Config{Resolvers: &backend.Resolver{DB: *db}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
