package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/rezaindrag/gographql"
)

type articleResolver struct {
}

func (r articleResolver) Fetch(params graphql.ResolveParams) (interface{}, error) {
	return "list of article", nil
}

// Initiator initiates the schema module.
type Initiator func(s *articleResolver) *articleResolver

func (i Initiator) Build() gographql.ArticleResolver {
	return i(&articleResolver{})
}

func NewSchema() Initiator {
	return func(s *articleResolver) *articleResolver {
		return s
	}
}
