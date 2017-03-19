package distribution

import (
	"github.com/gocarina/gocsv"
	"os"
)

type Country struct {
	Code string `csv:"Country Code"`
	Name string `csv:"Country Name`
}

type Province struct {
	Code string `csv:"Province Code"`
	Name string `csv:"Province Name"`
}

type City struct {
	Code string `csv:"City Code"`
	Name string `csv:"City Name"`
	Province
	Country
}

// Load the cities information from CSV file
func LoadCitiesFromCSV(filePath string) ([]*City, ApplicationError) {
	citiesFile, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, OSError(err.Error())
	}
	defer citiesFile.Close()

	cities := []*City{}
	if err := gocsv.UnmarshalFile(citiesFile, &cities); err != nil {
		return nil, InputError(err.Error())
	}
	return cities, nil
}

// Load the cities information from JSON file
func LoadCitiesFromJSON(filePath string) ([]*City, ApplicationError) {
	return nil, nil
}

// Load the cities information from XML file
func LoadCitiesFromXML(filePath string) ([]*City, ApplicationError) {
	return nil, nil
}
