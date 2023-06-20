package crssy

import "testing"

func TestFindCity(t *testing.T) {
	testdata := []struct {
		cityNames              []string
		wontCityCount          int
		wontFirstItemLatitude  float64
		wontFirstItemLongitude float64
	}{
		{[]string{"Kyoto"}, 1, 35.021070, 135.753850},
		{[]string{"Tokyo"}, 1, 35.689500, 139.691710},
		{[]string{"Unknown"}, 0, -1, -1},
	}
	for _, td := range testdata {
		results, _ := WeatherTemplate(td.cityNames)
		if len(results) != td.wontCityCount {
			t.Errorf("%v: wont %d city count, but got %d", td.cityNames, td.wontCityCount, len(results))
		}
		if len(results) > 0 {
			if results[0].Latitude != td.wontFirstItemLatitude && results[0].Longitude != td.wontFirstItemLongitude {
				t.Errorf("%s: wont (%f, %f), but got (%f, %f)", td.cityNames[0], td.wontFirstItemLatitude, td.wontFirstItemLongitude, results[0].Latitude, results[1].Longitude)
			}
		}
	}
}
