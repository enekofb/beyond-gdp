package countries

import (
	"encoding/csv"
	"os"

	"github.com/pkg/errors"
)

// Conf configuration entity for countries
type Conf struct {
	ResourcesPath string
}

//Country represents country entity in our domain
type Country struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

// CountryRepository repository pattern for country entity
type CountryRepository struct {
	items map[string]Country
}

func (r CountryRepository) GetAll() {

}

// NewRepository creates a new repository for country entities
func NewRepository(conf Conf) (CountryRepository, error) {
	countries, err := readCountriesFromCsv(conf.ResourcesPath)
	if err != nil {
		return CountryRepository{}, errors.Wrap(err, "cannot read countries from csv")
	}
	var countryMap = map[string]Country{}
	for _, country := range countries {
		countryMap[country.Name] = country
	}
	return CountryRepository{
		items: countryMap,
	}, nil
}

var countriesResource string
var countries []Country

func readCountriesFromCsv(csvPath string) ([]Country, error) {
	if csvPath == "" {
		return []Country{}, errors.New("invalid csv path")
	}
	exists, err := fileExists(csvPath)
	if err != nil || !exists {
		return []Country{}, errors.Wrap(err, "file does not exits")
	}
	countriesAsLines, err := readCsv(csvPath)
	if err != nil {
		return []Country{}, errors.Wrap(err, "cannot read csv file")
	}
	countries, err := adaptToCountries(countriesAsLines[1:])
	if err != nil {
		return []Country{}, errors.Wrap(err, "cannot adapt countries")
	}
	return countries, nil
}

func adaptToCountries(countriesAsLines [][]string) ([]Country, error) {
	countries := []Country{}
	for _, countryLines := range countriesAsLines {
		countries = append(countries, Country{
			Name:  countryLines[1],
			Score: countryLines[2],
		})
	}
	return countries, nil
}

func fileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false, errors.New("file does not exist")
	}
	return true, nil
}

func readCsv(csvPath string) ([][]string, error) {

	var results [][]string

	// open file
	f, err := os.Open(csvPath)
	if err != nil {
		return results, errors.Wrap(err, "cannot open file")
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	results, err = csvReader.ReadAll()
	if err != nil {
		return results, errors.Wrap(err, "cannot read all from csv file")
	}

	return results, nil
}

// Get the full list of countries
func (conf CountryRepository) Get() ([]Country, error) {
	if len(countries) == 0 {
		var err error
		countries, err = readCountriesFromCsv(conf.ResourcesPath)
		if err != nil {
			return []Country{}, errors.Wrap(err, "cannot read countries from csv")
		}
	}
	return countries, nil
}
