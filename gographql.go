package gographql

import (
	"github.com/graphql-go/graphql"
)

// ArticleResolver interface defines a list of contract methods.
type ArticleResolver interface {
	Fetch(params graphql.ResolveParams) (interface{}, error)
}
