package main

import (
	"fmt"
	"graphql-practice/router"
)

func main() {
	r := router.Router
	router.SetRouter()

	fmt.Println("Now server is running on port 8080")
	r.Run(":8080")

}
