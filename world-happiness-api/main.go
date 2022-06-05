package main

import (
	"encoding/json"
	"log"

	"github.com/enekofb/beyond-gdp/world-happiness-api/pkg/countries"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	resources := viper.GetString("country.resources")
	repository, err := countries.NewRepository(countries.Configuration{
		resources,
	})
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
	// setup setupConfiguration
	setupConfiguration(defaultConfigName)
	//setup router
	router := setupRouter()
	err := router.Run(":8080")
	if err != nil {
		log.Panic(errors.Wrap(err, "cannot run server"))
	}
}

//configuration file to be config.yaml within the app root directory
const defaultConfigName = "config"
const defaultConfigExtension = "yaml"

func setupConfiguration(configName string) {
	log.Printf("configuration started")
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType(defaultConfigExtension)
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Printf("cannot configure from file %s", err.Error())
		setupDefaultConfigurations()
	}
}

func setupDefaultConfigurations() {
	viper.SetDefault("country.resources", ".resources/world-happiness-data.csv")
}
