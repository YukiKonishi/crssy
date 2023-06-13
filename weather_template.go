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

func WeatherTemplate(args []string) int {
	cities := []*City{}
	err := json.Unmarshal(citiesJson, &cities)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	for _, city := range cities {
		city.Latitude, _ = strconv.ParseFloat(city.Lat, 64)
		city.Longitude, _ = strconv.ParseFloat(city.Lng, 64)
	}
	//件数をプリントしている
	//fmt.Printf("read %d entries\n", len(cities))
	for i := 0; i < len(cities); i++ {
		if args[1] == cities[i].Name {
			//一致した場所の国名，名前，緯度経度を返す
			fmt.Printf("Country: %s, Name: %s, Latitude: %f, Longitude: %f\n", cities[i].Country, cities[i].Name, cities[i].Latitude, cities[i].Longitude)
			MakeUrl(cities[i].Latitude, cities[i].Longitude)
		}
	}
	return 2
}
