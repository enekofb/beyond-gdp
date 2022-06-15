package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/enekofb/beyond-gdp/world-happiness-api/pkg/countries"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {

	resourcesAsCsv := viper.GetString("country.resources")
	repository, err := countries.NewRepositoryFromCsv(resourcesAsCsv)
	if err != nil {
		log.Panic(errors.Wrap(err, "cannot create country repository"))
		return "cannot create country repository", errors.Wrap(err, "cannot marshall json")
	}

	countries := repository.GetAll()
	countriesAsJson, err := json.Marshal(&countries)
	if err != nil {
		return "internal error", errors.Wrap(err, "cannot marshall json")
	}

	return fmt.Sprintf(string(countriesAsJson)), nil
	//c.String(http.StatusOK, string(countriesAsJson))
	//return
	//
	//
	//return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	// setup setupConfiguration
	setupConfiguration(defaultConfigName)

	lambda.Start(HandleRequest)
}

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
