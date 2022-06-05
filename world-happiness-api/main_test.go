package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const testConfigName = ".resources/test/config"

func TestConfiguration(t *testing.T) {

	t.Run("configurations with defaults", func(t *testing.T) {
		setupConfiguration("dontExits")
		assert.Equal(t, ".resources/world-happiness-data.csv", viper.GetString("country.resources"))
	})

	t.Run("can start with configuration", func(t *testing.T) {
		setupConfiguration(testConfigName)
		assert.Equal(t, ".resources/test/world-happiness-data.csv", viper.GetString("country.resources"))
	})

}

func TestRouter(t *testing.T) {

	t.Run("can check health endpoint", func(t *testing.T) {

		setupConfiguration(testConfigName)

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "up", w.Body.String())
	})

	t.Run("can get all countries", func(t *testing.T) {

		setupConfiguration(testConfigName)

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countries", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Finland")
		assert.Contains(t, w.Body.String(), "Denmark")

	})

	t.Run("can get country by countryName", func(t *testing.T) {
		var countryName = "Finland"
		setupConfiguration(testConfigName)

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/countries/%s", countryName), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Finland")
	})

}
