package schema

import (
	"graphql-practice/data"

	"github.com/graphql-go/graphql"
)

var getAuthor = &graphql.Field{
	// Type: graphql.String,
	Type: AuthorType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		if author, ok := p.Args["name"].(string); ok {
			for _, a := range data.AuthorList {
				if a.Name == author {
					return a, nil
				}
			}
		}
		return nil, nil
	},
}

var getAuthorAndBooks = &graphql.Field{
	Type: graphql.NewList(BookUnionType),
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		if author, ok := p.Args["name"].(string); ok {
			for _, a := range data.AuthorList {
				if a.Name == author {
					return a.Books, nil
				}
			}
		}
		return nil, nil
	},
}

var getBooks = &graphql.Field{
	Type: graphql.NewList(BookUnionType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return data.BookStore, nil
	},
}
