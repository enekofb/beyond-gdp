package countries

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCountriesRepository(t *testing.T) {
	conf := Conf{
		ResourcesPath: "../../.resources/world-happiness-data.csv",
	}

	repository, err := NewRepository(conf)

	require.Nil(t, err, "error not expected")
	require.NotNil(t, repository, "repository is nil")
}

func Test_readCountriesFromCsv(t *testing.T) {
	testCountriesResource := "../../.resources/world-happiness-data.csv"
	countries, err := readCountriesFromCsv(testCountriesResource)
	require.Nil(t, err, "error not expected")
	require.NotEmpty(t, countries, "should have some countries")
	require.Equal(t, len(countries), 147, "should have all countries")

}

func TestGetAll(t *testing.T) {
	conf := Conf{
		ResourcesPath: "../../.resources/world-happiness-data.csv",
	}
	repository, err := NewRepository(conf)
	require.Nil(t, err, "error not expected")
	countries := repository.GetAll()
	require.Equal(t, len(countries), 147, "should have all countries")
}
