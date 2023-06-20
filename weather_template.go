package crssy

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

//go:embed "cities.json"
var citiesJson []byte

type City struct {
	Country   string  `json:"country"`
	Name      string  `json:"name"`
	Lat       string  `json:"lat"`
	Lng       string  `json:"lng"`
	Latitude  float64 `json:"-"`
	Longitude float64 `json:"-"`
}

func findCity(cityName string, cities []*City) (*City, error) {
	//件数をプリントしている
	//fmt.Printf("read %d entries\n", len(cities))
	for i := 0; i < len(cities); i++ {
		if cityName == cities[i].Name {
			return cities[i], nil
		}
	}
	return nil, fmt.Errorf("%s: city not found", cityName)
}

func WeatherTemplate(args []string) ([]*City, error) {
	cities := []*City{}
	err := json.Unmarshal(citiesJson, &cities)
	if err != nil {
		return nil, err
	}
	fmt.Printf("load %d cities\n", len(cities))
	for _, city := range cities {
		city.Latitude, _ = strconv.ParseFloat(city.Lat, 64)
		city.Longitude, _ = strconv.ParseFloat(city.Lng, 64)
	}
	var results []*City
	var errs []error
	for _, name := range args {
		city, err := findCity(name, cities)
		if err == nil {
			results = append(results, city)
		} else {
			errs = append(errs, err)
		}
	}
	return results, errors.Join(errs...)
}
