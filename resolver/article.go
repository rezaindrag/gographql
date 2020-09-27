package resolver

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/rezaindrag/gographql"
	"net/http"
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
	id, ok := params.Args["id"].(string)
	if !ok || id == "" {
		return gographql.Article{}, gographql.ExtendedError{
			Err:        errors.New("id is not provided"),
			StatusCode: http.StatusBadRequest,
		}
	}

	article := gographql.Article{
		ID:    id,
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
