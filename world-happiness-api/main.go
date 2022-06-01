package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "up")
	})

	return router
}

func main() {

	router := setupRouter()
	router.Run(":8080")
}
