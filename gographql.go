package gographql

import (
	"github.com/graphql-go/graphql"
)

// ArticleResolver interface defines a list of contract methods.
type ArticleResolver interface {
	Fetch(params graphql.ResolveParams) (interface{}, error)
	GetByID(params graphql.ResolveParams) (interface{}, error)
	Create(params graphql.ResolveParams) (interface{}, error)
	Update(params graphql.ResolveParams) (interface{}, error)
}
