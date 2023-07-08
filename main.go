package main

import (
	"github.com/adpadillar/htmx-go/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	router.Static("/public", "./src/public")

	router.GET("/", routes.IndexHandler)
	router.GET("/hello", routes.HelloHandler)
	router.GET("/new-page", routes.NewPageHandler)

	// Start the server
	router.Run(":8081")
}
