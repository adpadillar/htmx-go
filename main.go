package main

import (
	"github.com/adpadillar/htmx-go/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	router.GET("/", routes.IndexHandler)
	router.GET("/hello", routes.HelloHandler)

	// Start the server
	router.Run(":8081")
}
