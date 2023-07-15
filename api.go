package crssy

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "net/http"
)

//go:embed "weather.json"
var weatherJson []byte

type Weather struct {
	Code string `json:"code"`
	Jp   string `json:"jp"`
	En   string `json:"en"`
}

type Result struct {
	Day *Daily `json:"daily"`
}

type Daily struct {
	Weathercode []int    `json:"weathercode`
	Time        []string `json:"time`
}

func ExpectWeather(city *City) (*Daily, error) {
	code, err := GetWeatherCode(city.Lat, city.Lng)
	if err != nil {
		return nil, err
	}
	return code, nil
}

func Translateweather(code int) (string, error) {
	weathers := []*Weather{}
	err := json.Unmarshal(weatherJson, &weathers)
	if err != nil {
		return "", err
	}
	// fmt.Println(weathers[code].Jp)
	return weathers[code].Jp, err
}

func GetWeatherCode(lat, lon string) (*Daily, error) {
	// fmt.Printf("%s,%s\n", lat, lon)
	timezone := "Asia/Tokyo"

	// URLの作成
	// url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=weathercode&forecast_days=1", lat, lon)
	// url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=weathercode", lat, lon)
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&daily=weathercode&timezone=%s&forecast_days=1", lat, lon, timezone)
	// fmt.Println(url)

	// HTTPリクエストの作成
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return nil, err
	}

	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの送信
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました:", err)
		return nil, err
	}
	defer response.Body.Close()

	// レスポンスの読み取り
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("レスポンスの読み取りに失敗しました:", err)
		return nil, err
	}

	result := Result{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	// fmt.Println(result.Day.Weathercode)

	return result.Day, nil

	// レスポンスの表示
	// responseString := string(body)

	//fmt.Println(responseString)
	// responseParts := strings.Split(responseString, ",")
	// lastIndex := len(responseParts) - 1
	// for i, part := range responseParts {
	// 	if i == lastIndex {
	// 		fmt.Println(part)
	// 		return part, nil
	// 	}
	// }
	// return 0, fmt.Errorf("コードが見つかりませんでした：%s", body)
}

//https://api.open-meteo.com/v1/forecast?latitude=35.02&longitude=135.75&hourly=weathercode&forecast_days=1
