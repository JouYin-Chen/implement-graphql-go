package router

import (
	"graphql-practice/graphqlFunc"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {

	Router.POST("/graphql", graphqlFunc.GraphqlHandler())
	Router.GET("/graphql", graphqlFunc.GraphqlHandler())
}
