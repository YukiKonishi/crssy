package crssy

import (
	_ "embed"
	"encoding/json"
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

var cities []*City

func FindCity(cityName string) (*City, error) {
	if len(cities) == 0 {
		if err := loadCities(); err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(cities); i++ {
		if cityName == cities[i].Name {
			// lat, _ := strconv.ParseFloat(cities[i].Lat, 64)
			// lng, _ := strconv.ParseFloat(cities[i].Lng, 64)
			// MakeUrl(cities[i].Lat, cities[i].Lng)
			return cities[i], nil
		}
	}
	return nil, fmt.Errorf("%s: city not found", cityName)
}

func loadCities() error {
	err := json.Unmarshal(citiesJson, &cities)
	if err != nil {
		return err
	}
	//fmt.Printf("load %d cities\n", len(cities))
	for _, city := range cities {
		city.Latitude, _ = strconv.ParseFloat(city.Lat, 64)
		city.Longitude, _ = strconv.ParseFloat(city.Lng, 64)
	}
	return nil
}
