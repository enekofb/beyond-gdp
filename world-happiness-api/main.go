package main

import (
	"encoding/json"
	"log"

	"github.com/enekofb/beyond-gdp/world-happiness-api/pkg/countries"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	countriesConf := countries.Conf{
		ResourcesPath: ".resources/world-happiness-data.csv",
	}

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "up")
	})

	router.GET("/countries", func(c *gin.Context) {
		countries, err := countriesConf.Get()
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot get all countries"))
			c.String(500, "internal error")
			return
		}
		countriesAsJson, err := json.Marshal(&countries)
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot marshall json"))
			c.String(500, "internal error")
		}
		c.String(200, string(countriesAsJson))
		return

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
