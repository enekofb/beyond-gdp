package countries

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

	countries, err := conf.Get()

	require.NotEmpty(t, countries, "should have some countries")
	require.Nil(t, err, "error not expected")
	require.Equal(t, len(countries), 147, "should have all countries")

}
