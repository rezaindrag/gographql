package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/rezaindrag/gographql"
)

type Schema struct {
	articleResolver gographql.ArticleResolver
}

func (s Schema) Query() *graphql.Object {
	fields := graphql.Fields{
		"articles": &graphql.Field{
			Type:        graphql.NewList(Article),
			Description: "Get a list of articles.",
			Resolve:     s.articleResolver.Fetch,
		},
		"article": &graphql.Field{
			Type:        Article,
			Description: "Get detail of article.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: s.articleResolver.GetByID,
		},
	}

	query := graphql.ObjectConfig{
		Name:   "Query",
		Fields: fields,
	}

	return graphql.NewObject(query)
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
