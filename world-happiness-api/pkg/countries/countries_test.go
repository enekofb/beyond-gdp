package countries

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCountriesRepository(t *testing.T) {
	csvPath := "../../.resources/test/world-happiness-data.csv"
	repository, err := NewRepositoryFromCsv(csvPath)

	require.Nil(t, err, "error not expected")
	require.NotNil(t, repository, "repository is nil")
}

func Test_readCountriesFromCsv(t *testing.T) {
	testCountriesResource := "../../.resources/test/world-happiness-data.csv"

	countries, err := readCountriesFromCsv(testCountriesResource)

	require.Nil(t, err, "error not expected")
	require.NotEmpty(t, countries, "should have some countries")
	require.Equal(t, len(countries), 2, "should have all countries")
}

func TestGetAll(t *testing.T) {
	csvPath := "../../.resources/test/world-happiness-data.csv"
	repository, err := NewRepositoryFromCsv(csvPath)
	require.Nil(t, err, "error not expected")

	countries := repository.GetAll()

	require.Equal(t, len(countries), 2, "should have all countries")
}

func TestGetByNameForExistingCountry(t *testing.T) {
	csvPath := "../../.resources/test/world-happiness-data.csv"
	repository, err := NewRepositoryFromCsv(csvPath)
	require.Nil(t, err, "error not expected")

	country, err := repository.GetByName("Finland")

	require.Nil(t, err, "error not expected")
	require.Equal(t, country.Name, "Finland", "not valid retrieved country")
}

func TestGetByNameForNonExistingCountry(t *testing.T) {
	csvPath := "../../.resources/test/world-happiness-data.csv"
	repository, err := NewRepositoryFromCsv(csvPath)
	require.Nil(t, err, "error not expected")

	country, err := repository.GetByName("dontExists")

	require.Nil(t, err, "error not expected")
	require.Equal(t, country, Country{}, "not valid retrieved country")
}
