package resolver

import (
	"encoding/json"
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
			Err:        errors.New("parameter id is not provided"),
			StatusCode: http.StatusBadRequest,
		}
	}

	article := gographql.Article{
		ID:    id,
		Title: "Foo",
	}

	return article, nil
}

func (r articleResolver) Create(params graphql.ResolveParams) (interface{}, error) {
	strArticle, ok := params.Args["article"].(string)
	if !ok || strArticle == "" {
		return gographql.Article{}, gographql.ExtendedError{
			Err:        errors.New("parameter article is not provided"),
			StatusCode: http.StatusBadRequest,
		}
	}

	newArticle := gographql.Article{}
	if err := json.Unmarshal([]byte(strArticle), &newArticle); err != nil {
		return gographql.Article{}, gographql.ExtendedError{
			Err:        errors.New("parameter article is invalid"),
			StatusCode: http.StatusBadRequest,
		}
	}

	article := newArticle

	return article, nil
}

func (r articleResolver) Update(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(string)
	if id == "" || !ok {
		return gographql.Article{}, gographql.ExtendedError{
			Err:        errors.New("parameter id is not provided"),
			StatusCode: http.StatusBadRequest,
		}
	}

	strArticle, ok := params.Args["article"].(string)
	if !ok || strArticle == "" {
		return gographql.Article{}, gographql.ExtendedError{
			Err:        errors.New("parameter article is not provided"),
			StatusCode: http.StatusBadRequest,
		}
	}

	modifiedArticle := gographql.Article{}
	if err := json.Unmarshal([]byte(strArticle), &modifiedArticle); err != nil {
		return gographql.Article{}, gographql.ExtendedError{
			Err:        errors.New("parameter article is invalid"),
			StatusCode: http.StatusBadRequest,
		}
	}

	article := gographql.Article{
		ID:    "foo",
		Title: modifiedArticle.ID,
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
