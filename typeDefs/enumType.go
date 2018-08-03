package typeDefs

import (
	"github.com/graphql-go/graphql"
)

var AuthorStateEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name: "AuthorStateEnum",
	Values: graphql.EnumValueConfigMap{
		"female": &graphql.EnumValueConfig{
			Value: 0,
		},
		"male": &graphql.EnumValueConfig{
			Value: 1,
		},
		"other": &graphql.EnumValueConfig{
			Value: 2,
		},
	},
})
