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

	repository, err := countries.NewRepository(countriesConf)
	if err != nil {
		log.Panic(errors.Wrap(err, "cannot create country repository"))
		return nil
	}

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "up")
	})

	router.GET("/countries", func(c *gin.Context) {
		countries := repository.GetAll()
		countriesAsJson, err := json.Marshal(&countries)
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot marshall json"))
			c.String(500, "internal error")
		}
		c.String(200, string(countriesAsJson))
		return
	})

	router.GET("/countries/:countryName", func(c *gin.Context) {
		log.Println("in country name")
		countryName := c.Param("countryName")
		country, err := repository.GetByName(countryName)
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot get country by name"))
			c.String(500, "internal error")
			return
		}
		var emptyCountry = countries.Country{}
		if country == emptyCountry {
			c.String(404, "country not found")
			return
		}
		countryAsJson, err := json.Marshal(&country)
		if err != nil {
			log.Fatal(errors.Wrap(err, "cannot marshall json"))
			c.String(500, "internal error")
		}
		c.String(200, string(countryAsJson))
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
