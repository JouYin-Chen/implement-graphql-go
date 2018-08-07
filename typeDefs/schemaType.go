package typeDefs

import (
	"graphql-practice/data"

	"github.com/graphql-go/graphql"
)

// described these types as graphql objects
var TextBookType *graphql.Object
var ShortStoryType *graphql.Object

func init() {
	TextBookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "TextBook",
		Interfaces: []*graphql.Interface{
			BookInterface,
		},
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"series": &graphql.Field{
				Type: graphql.Int,
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(*data.TextBook)
			return ok
		},
	})

	ShortStoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "ShortStory",
		Interfaces: []*graphql.Interface{
			BookInterface,
		},
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"chapters": &graphql.Field{
				Type: graphql.Int,
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(*data.ShortStory)
			return ok
		},
	})

}
