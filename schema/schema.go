package schema

import (
	"graphql-practice/data"
	"graphql-practice/typeDefs"

	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

var BookUnionType = graphql.NewUnion(graphql.UnionConfig{
	Name: "BookList",
	Types: []*graphql.Object{
		typeDefs.TextBookType, typeDefs.ShortStoryType,
	},
	ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
		if _, ok := p.Value.(*data.TextBook); ok {
			return typeDefs.TextBookType
		}
		if _, ok := p.Value.(*data.ShortStory); ok {
			return typeDefs.ShortStoryType
		}
		return nil
	},
})

var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"name":   &graphql.Field{Type: graphql.String},
		"age":    &graphql.Field{Type: graphql.Int},
		"books":  &graphql.Field{Type: graphql.NewList(BookUnionType)},
		"gendor": {Type: typeDefs.AuthorStateEnumType},
	},
})

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
		"getAuthor":         getAuthor,
		"getAuthorAndBooks": getAuthorAndBooks,
		"getBooks":          getBooks,
	},
	// IsTypeOf: func(p graphql.IsTypeOfParams) bool {
	// 	_, ok := p.Value.(*data.Author)
	// 	return ok
	// },
})

// "getBooksByTitle": &graphql.Field{
// 	Type: graphql.NewObject(graphql.ObjectConfig{
// 		Name: "getBookByTitleInterface",
// 		Fields: graphql.Fields{
// 			"title": {
// 				Type: graphql.String,
// 			},
// 			"series": {
// 				Type: graphql.Int,
// 			},
// 		},
// 		Interfaces: []*graphql.Interface{
// 			typeDefs.BookInterface,
// 		},
// 	}),
// 	Args: graphql.FieldConfigArgument{
// 		"title": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		if title, ok := p.Args["title"].(string); ok {
// 			for _, d := range data.TextBookList {
// 				if title == d.Title {
// 					return d, nil
// 				}
// 			}
// 		}
// 		return nil, nil
// 	},
// },
// "getBookAndAuthor": &graphql.Field{
// 	Type: graphql.NewObject(graphql.ObjectConfig{
// 		Name: "getBookAndAuthor",
// 		Fields: graphql.Fields{
// 			"author": {
// 				Type: typeDefs.AuthorType,
// 			},
// 			"book": {
// 				Type: typeDefs.TextBookType,
// 			},
// 		},
// 	}),
// 	Args: graphql.FieldConfigArgument{
// 		"title": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 	},
// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 		if title, ok := p.Args["title"].(string); ok {
// 			for _, d := range data.TextBookList {
// 				if d.Title == title {
// 					for _, a := range data.AuthorList {
// 						if d.Author == a.Name {
// 							return struct {
// 								Author data.Author
// 								Book   data.TextBook
// 							}{a, d}, nil
// 						}
// 					}
// 					return d, nil
// 				}
// 			}
// 		}
// 		return nil, nil
// 	},
// },
// 	},
// })

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createAuthor": &graphql.Field{
			Type: graphql.NewList(AuthorType), // the return type for this field

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// marshall and cast the argument value
				// perform mutation operation here
				// for e.g. create a Todo and save to DB.

				newAuthor := data.Author{
					Name:   "Louis Cha Leung-yung",
					Age:    94,
					Gendor: 1,
					Books: []data.BookList{
						data.TextBook{"射鵰英雄傳", 4},
						data.TextBook{"神鵰俠侶", 4},
						data.TextBook{"倚天屠龍記", 4},
					},
				}

				data.AuthorList = append(data.AuthorList, newAuthor)
				// return the new Todo object that we supposedly save to DB
				// Note here that
				// - we are returning a `Todo` struct instance here
				// - we previously specified the return Type to be `todoType`
				// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
				return data.AuthorList, nil
			},
		},
	},
})
