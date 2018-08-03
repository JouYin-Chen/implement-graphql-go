package graphqlFunc

import (
	"graphql-practice/schema"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Gin封装
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
