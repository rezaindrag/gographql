package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
	"github.com/rezaindrag/gographql/resolver"
	"github.com/rezaindrag/gographql/schema"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	articleResolver := resolver.NewArticleResolver().Build()

	mySchema := schema.NewSchema().WithArticleResolver(articleResolver).Build()

	gqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: mySchema.Query(),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	e.POST("/graphql", echo.WrapHandler(handler.New(&handler.Config{
		Schema: &gqlSchema,
	})))

	graphiqlHandler := echo.WrapHandler(handler.New(&handler.Config{
		Schema:   &gqlSchema,
		GraphiQL: true,
	}))
	e.GET("/graphiql", graphiqlHandler)
	e.POST("/graphiql", graphiqlHandler)

	if err := e.Start(":7723"); err != nil {
		logrus.Fatal(err)
	}
}
