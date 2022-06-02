package main

import (
	"log"

	"github.com/enekofb/beyond-gdp/world-happiness-api/pkg/countries"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "up")
	})

	router.GET("/countries", func(c *gin.Context) {
		_, err := countries.GetAll()
		if err != nil {
			log.Println("cannot get all countries")
			c.String(500, "not available")
			return
		}
	})

	return router
}

func main() {

	router := setupRouter()
	err := router.Run(":8080")
	//todo: review
	if err != nil {
		log.Panic(err)
	}
}
