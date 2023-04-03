package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JoseAngel1196/weather/config"
	"github.com/JoseAngel1196/weather/share"
)

const (
	geocodeApiUrl = "https://maps.googleapis.com/maps/api/geocode/json"
)

func GetGeolocation(address share.UserAddress) (share.Location, error) {
	var location share.Location

	googleApiKey, err := config.GetEnv("GOOGLE_API_KEY")
	if err != nil {
		return location, err
	}

	formattedAddress := address.FormatAddress()
	formattedAddress = strings.Replace(formattedAddress, " ", "+", -1)
	httpUrl := fmt.Sprintf("%v?key=%s&address=%s", geocodeApiUrl, googleApiKey, formattedAddress)

	body, err := httpRequest(httpUrl)
	if err != nil {
		return location, err
	}

	var geoResults share.GeoResults
	resp, err := unmarshalResponse(body)
	if err != nil {
		return location, err
	}

	geoResults, ok := resp.(share.GeoResults)
	if !ok {
		return location, errors.New("Cannot convert ResponseData to GeoResults")
	}

	if strings.ToUpper(geoResults.Status) != "OK" {
		switch strings.ToUpper(geoResults.Status) {
		case "ZERO_RESULTS":
			err = errors.New("No results found.")
			break
		case "OVER_QUERY_LIMIT":
			err = errors.New("You are over your quota.")
			break
		case "REQUEST_DENIED":
			err = errors.New("Your request was denied.")
			break
		case "INVALID_REQUEST":
			err = errors.New("Probably the query is missing.")
			break
		case "UNKNOWN_ERROR":
			err = errors.New("Server error. Please, try again.")
			break
		default:
			break
		}
	}

	location.Latitude = geoResults.Results[0].Geometry.Location.Lat
	location.Longitude = geoResults.Results[0].Geometry.Location.Lng

	return location, nil
}
