package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/JoseAngel1196/weather/share"
)

type ResponseData interface {
}

func httpRequest(url string) ([]byte, error) {
	// create a new http client
	client := &http.Client{}

	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// make sure the response body gets closed after this function is done
	defer resp.Body.Close()

	// read the response body into a byte array
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// return the response body
	return body, nil
}

func unmarshalResponse(body []byte) (ResponseData, error) {
	// define a variable to hold the response struct
	var response ResponseData

	// unmarshal the response body into the appropriate struct based on the API response
	var dataA share.GeoResults
	if json.Unmarshal(body, &dataA) == nil && dataA.Results != nil {
		response = dataA
	} else {
		var dataB share.WeatherPoints
		if json.Unmarshal(body, &dataB) == nil && dataB.ID != "" {
			response = dataB
		} else {
			// if the response is not of a known type, return an error
			return nil, errors.New("unknown response type")
		}
	}

	return response, nil
}
