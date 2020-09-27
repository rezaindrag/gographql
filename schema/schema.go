package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/rezaindrag/gographql"
)

type Schema struct {
	articleResolver gographql.ArticleResolver
}

func (s Schema) Query() *graphql.Object {
	articleField := graphql.Field{
		Name:              "",
		Type:              graphql.String,
		Args:              nil,
		Resolve:           s.articleResolver.Fetch,
		DeprecationReason: "",
		Description:       "",
	}

	fields := graphql.Fields{
		"articles": &articleField,
	}

	query := graphql.ObjectConfig{
		Name:        "",
		Interfaces:  nil,
		Fields:      fields,
		IsTypeOf:    nil,
		Description: "",
	}

	return graphql.NewObject(query)
}

func (s Schema) Mutation() *graphql.Object {
	mutation := graphql.ObjectConfig{
		Name:        "",
		Interfaces:  nil,
		Fields:      nil,
		IsTypeOf:    nil,
		Description: "",
	}
	return graphql.NewObject(mutation)
}

// Initiator initiates the schema module.
type Initiator func(s *Schema) *Schema

// WithArticleResolver initiates with the article resolver.
func (i Initiator) WithArticleResolver(articleResolver gographql.ArticleResolver) Initiator {
	return func(s *Schema) *Schema {
		i(s).articleResolver = articleResolver
		return s
	}
}

func (i Initiator) Build() *Schema {
	return i(&Schema{})
}

func NewSchema() Initiator {
	return func(s *Schema) *Schema {
		return s
	}
}
