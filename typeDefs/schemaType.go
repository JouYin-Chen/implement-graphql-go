package typeDefs

import (
	"github.com/graphql-go/graphql"
)

// described these types as graphql objects
var TextBookType *graphql.Object

var AuthorType *graphql.Object

func init() {
	TextBookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "TextBook",
		Interfaces: []*graphql.Interface{
			BookInterface,
		},
		Fields: graphql.Fields{
			"title":  &graphql.Field{Type: graphql.String},
			"series": &graphql.Field{Type: graphql.Int},
			"author": &graphql.Field{Type: graphql.String},
		},
	})

	AuthorType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"name":  &graphql.Field{Type: graphql.String},
			"age":   &graphql.Field{Type: graphql.Int},
			"books": &graphql.Field{Type: graphql.NewList(TextBookType)},
		},
	})

}
