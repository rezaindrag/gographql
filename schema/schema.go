package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/rezaindrag/gographql"
)

type Schema struct {
	articleResolver gographql.ArticleResolver
}

var json = graphql.NewScalar(
	graphql.ScalarConfig{
		Name: "JSON",
		ParseValue: func(value interface{}) interface{} {
			if value == nil {
				return make(map[string]interface{})
			}
			return value
		},
		Description: "The JSON scalar type represents JavaScript object notation syntax.",
	},
)

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

func (s Schema) Mutation() *graphql.Object {
	fields := graphql.Fields{
		"createArticle": &graphql.Field{
			Type:        Article,
			Description: "Create a new article.",
			Args: graphql.FieldConfigArgument{
				"article": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(json),
				},
			},
			Resolve: s.articleResolver.Create,
		},
		"updateArticle": &graphql.Field{
			Type:        Article,
			Description: "Update an article.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"article": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(json),
				},
			},
			Resolve: s.articleResolver.Update,
		},
	}

	query := graphql.ObjectConfig{
		Name:   "Mutation",
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
