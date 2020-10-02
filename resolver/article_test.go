package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/pkg/errors"
	"github.com/rezaindrag/gographql"
	"github.com/stretchr/testify/assert"
	"net/http"
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
					"id": "foo",
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
				ID:    "foo",
				Title: "Foo",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			articleResolver := NewArticleResolver().Build()
			result, err := articleResolver.GetByID(test.passParams)
			if err != nil {
				assert.EqualError(t, errors.Cause(err), test.exptError.Err.Error())
				return
			}
			assert.Equal(t, result, test.exptResult)
		})
	}
}

func TestArticleResolver_Create(t *testing.T) {
	for name, test := range map[string]struct {
		passParams graphql.ResolveParams
		exptResult interface{}
		exptError  gographql.ExtendedError
	}{
		"Successfully resolve create new article": {
			passParams: graphql.ResolveParams{
				Args: map[string]interface{}{
					"article": "{\"id\":\"bar\",\"title\":\"Bar\"}",
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
				ID:    "bar",
				Title: "Bar",
			},
		},
		"Error parameter article is not provided": {
			passParams: graphql.ResolveParams{
				Args: map[string]interface{}{},
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
			exptError: gographql.ExtendedError{
				Err:        errors.New("parameter article is not provided"),
				StatusCode: http.StatusBadRequest,
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			articleResolver := NewArticleResolver().Build()
			result, err := articleResolver.Create(test.passParams)
			if err != nil {
				assert.EqualError(t, errors.Cause(err), test.exptError.Err.Error())
				return
			}
			assert.Equal(t, result, test.exptResult)
		})
	}
}
