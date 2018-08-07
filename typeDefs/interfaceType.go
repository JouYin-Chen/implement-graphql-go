package typeDefs

import (
	"github.com/graphql-go/graphql"
)

var BookInterface = graphql.NewInterface(graphql.InterfaceConfig{
	Name: "BookInterface",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})
