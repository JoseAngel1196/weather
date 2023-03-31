package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/JoseAngel1196/weather/config"
)

type WeatherDataResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func GetWeatherData(userLocation userLocation) (WeatherDataResponse, error) {
	openWeatherApiKey, err := config.GetEnv("OPEN_WEATHER_API_KEY")
	if err != nil {
		return WeatherDataResponse{}, err
	}

	latitude := userLocation.latitude
	longitude := userLocation.longitude
	httpUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.2f&lon=%.2f&appid=%s", latitude, longitude, url.QueryEscape(openWeatherApiKey))

	resp, err := http.Get(httpUrl)
	if err != nil {
		return WeatherDataResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll((resp.Body))
	if err != nil {
		return WeatherDataResponse{}, err
	}

	var bodyAPI WeatherDataResponse
	err = json.Unmarshal(body, &bodyAPI)
	if err != nil {
		return WeatherDataResponse{}, err
	}

	return bodyAPI, nil
}
