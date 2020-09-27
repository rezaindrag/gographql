package schema

import "github.com/graphql-go/graphql"

var Article = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
