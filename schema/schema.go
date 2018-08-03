package schema

import (
	"graphql-practice/data"
	"graphql-practice/typeDefs"

	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)

/*
   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getAuthorByName": &graphql.Field{
			Type: typeDefs.AuthorType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if author, ok := p.Args["name"].(string); ok {
					for _, d := range data.AuthorList {
						if d.Name == author {
							return d, nil
						}
					}
				}
				return nil, nil
			},
		},
		"getBooks": &graphql.Field{
			Type: graphql.NewList(typeDefs.TextBookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return data.TextBookList, nil
			},
		},
		"getBookAndAuthor": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "getBookAndAuthor",
				Fields: graphql.Fields{
					"author": {
						Type: typeDefs.AuthorType,
					},
					"book": {
						Type: typeDefs.TextBookType,
					},
				},
			}),
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				if title, ok := p.Args["title"].(string); ok {
					for _, d := range data.TextBookList {
						if d.Title == title {
							for _, a := range data.AuthorList {
								if d.Author == a.Name {
									return struct {
										Author data.Author
										Book   data.TextBook
									}{a, d}, nil
								}
							}
							return d, nil
						}
					}
				}
				return nil, nil
			},
		},
	},
})
