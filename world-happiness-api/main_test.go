package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter_Health(t *testing.T) {

	t.Run("can check health endpoint", func(t *testing.T) {

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "up", w.Body.String())
	})

	t.Run("can get all countries", func(t *testing.T) {

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countries", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Finland")
		assert.Contains(t, w.Body.String(), "Spain")

	})

	t.Run("can get country by countryName", func(t *testing.T) {
		var countryName = "Spain"

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/countries/%s", countryName), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Spain")
	})

}
