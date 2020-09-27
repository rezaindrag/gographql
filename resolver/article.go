package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/rezaindrag/gographql"
)

type articleResolver struct {
}

func (r articleResolver) Fetch(params graphql.ResolveParams) (interface{}, error) {
	articles := []gographql.Article{
		{
			ID:    "foo",
			Title: "Foo",
		},
	}

	return articles, nil
}

func (r articleResolver) GetByID(params graphql.ResolveParams) (interface{}, error) {
	article := gographql.Article{
		ID:    "foo",
		Title: "Foo",
	}

	return article, nil
}

// Initiator initiates the schema module.
type Initiator func(s *articleResolver) *articleResolver

func (i Initiator) Build() gographql.ArticleResolver {
	return i(&articleResolver{})
}

func NewArticleResolver() Initiator {
	return func(s *articleResolver) *articleResolver {
		return s
	}
}
