package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfiguration(t *testing.T) {

	t.Run("can start with default setupConfiguration", func(t *testing.T) {
		err := setupConfiguration()
		require.Nil(t, err)

	})

}

func TestRouter(t *testing.T) {

	t.Run("can check health endpoint", func(t *testing.T) {

		err := setupConfiguration()
		require.Nil(t, err)

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "up", w.Body.String())
	})

	t.Run("can get all countries", func(t *testing.T) {

		err := setupConfiguration()
		require.Nil(t, err)

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
		err := setupConfiguration()
		require.Nil(t, err)

		router := setupRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/countries/%s", countryName), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Spain")
	})

}
