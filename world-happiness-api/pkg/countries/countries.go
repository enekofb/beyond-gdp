package countries

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/pkg/errors"
)

//Country represents country entity in our domain
type Country struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

// CountryRepository repository pattern for country entity
type CountryRepository struct {
	asList []Country
	asMap  map[string]Country
}

func (repository CountryRepository) GetAll() []Country {
	return repository.asList
}

func (repository CountryRepository) GetByName(name string) (Country, error) {
	if name == "" {
		return Country{}, errors.New("country name cannot be empty")
	}
	if country, ok := repository.asMap[name]; ok {
		return country, nil
	}
	return Country{}, nil
}

// NewRepositoryFromCsv creates a new repository for country entities using csv as data source
func NewRepositoryFromCsv(csvPath string) (CountryRepository, error) {
	if csvPath == "" {
		return CountryRepository{}, errors.New("csv path cannot be empty")
	}
	log.Printf("create repository from csv file in '%s'", csvPath)
	countries, err := readCountriesFromCsv(csvPath)
	if err != nil {
		return CountryRepository{}, errors.Wrap(err, "cannot read countries from csv")
	}
	var countryMap = map[string]Country{}
	for _, country := range countries {
		countryMap[country.Name] = country
	}
	return CountryRepository{
		asMap:  countryMap,
		asList: countries,
	}, nil
}

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
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf(errors.Wrap(err, "cannot close csv file").Error())
		}
	}(f)

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	results, err = csvReader.ReadAll()
	if err != nil {
		return results, errors.Wrap(err, "cannot read all from csv file")
	}

	return results, nil
}
