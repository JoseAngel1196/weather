package api

import (
	"fmt"

	"github.com/JoseAngel1196/weather/share"
)

const (
	weatherGovUrl = "https://api.weather.gov"
)

func GetWeatherPoints(latitude, longitude float64) (share.WeatherPoints, error) {
	httpUrl := fmt.Sprintf("%v/points/%.4f,%.4f", weatherGovUrl, latitude, longitude)
	var weatherPoints share.WeatherPoints

	body, err := httpRequest(httpUrl)
	if err != nil {
		return weatherPoints, err
	}

	err = unmarshalResponse(body, &weatherPoints)
	if err != nil {
		return weatherPoints, err
	}

	return weatherPoints, nil
}

func GetForecastHourly(forecastHourlyUrl string) (share.ForecastHourly, error) {
	var forecastHourly share.ForecastHourly

	body, err := httpRequest(forecastHourlyUrl)
	if err != nil {
		return forecastHourly, err
	}

	err = unmarshalResponse(body, &forecastHourly)
	if err != nil {
		return forecastHourly, err
	}

	return forecastHourly, nil
}
