package api

import (
	"errors"
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

	resp, err := unmarshalResponse(body)
	if err != nil {
		return weatherPoints, err
	}

	weatherPoints, ok := resp.(share.WeatherPoints)
	if !ok {
		return weatherPoints, errors.New("Cannot convert ResponseData to PointsData")
	}

	return weatherPoints, nil
}
