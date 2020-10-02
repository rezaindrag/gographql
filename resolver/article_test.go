package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/pkg/errors"
	"github.com/rezaindrag/gographql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArticleResolver_GetByID(t *testing.T) {
	for name, test := range map[string]struct {
		passParams graphql.ResolveParams
		exptResult interface{}
		exptError  gographql.ExtendedError
	}{
		"Successfully resolve getting article by ID": {
			passParams: graphql.ResolveParams{
				Args: map[string]interface{}{
					"id": "2",
				},
				Info: graphql.ResolveInfo{
					FieldASTs: []*ast.Field{
						{
							SelectionSet: ast.NewSelectionSet(&ast.SelectionSet{
								Selections: []ast.Selection{
									ast.NewField(&ast.Field{
										Name: ast.NewName(&ast.Name{
											Value: "id",
										}),
									}),
									ast.NewField(&ast.Field{
										Name: ast.NewName(&ast.Name{
											Value: "title",
										}),
									}),
								},
							}),
						},
					},
				},
			},
			exptResult: gographql.Article{
				ID:    "2",
				Title: "Foo",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			articleResolver := NewArticleResolver().Build()
			result, err := articleResolver.GetByID(test.passParams)
			if err != nil {
				assert.EqualError(t, errors.Cause(err), test.exptError.Err.Error())
			}
			assert.Equal(t, result, test.exptResult)
		})
	}

}
