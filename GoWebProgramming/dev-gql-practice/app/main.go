package main

import (
	"app/graph"
	"app/graph/generated"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	// "github.com/99designs/gqlgen/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "host=db user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		log.Fatalln("接続失敗", err)
	} else {
		fmt.Println("--init--")
	}
	defer db.Close()
	db.LogMode(true)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: *db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// e.POST("/query", func(c echo.Context) error {
	// 	config := generated.Config{
	// 		Resolvers: &graph.Resolver{},
	// 	}
	// 	h := handler.GraphQL(generated.NewExecutableSchema(config))
	// 	h.ServeHTTP(c.Response(), c.Request())
	// 	return nil
	// })

	// e.Logger.SetLevel(elog.INFO)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":8080"))
}
