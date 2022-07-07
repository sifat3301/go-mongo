package main

import (
	"github.com/gin-gonic/gin"
	"go-mongo/configs"
	"go-mongo/routes"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})

	})
	// Run database
	configs.ConnectDB()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	//routes
	routes.RootRoute(router)

	router.Run(host + ":" + port)
}
