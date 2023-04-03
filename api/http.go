package api

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
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

func unmarshalResponse(body []byte, v ResponseData) error {
	// create a new value of the same type as the second argument
	// this will be used to unmarshal the response into the correct struct
	valueType := reflect.ValueOf(v).Elem().Type()
	valuePtr := reflect.New(valueType)

	// use the reflect package to determine the type of the response
	err := json.Unmarshal(body, valuePtr.Interface())
	if err != nil {
		return err
	}

	// update the value of the v argument with the unmarshaled value
	reflect.ValueOf(v).Elem().Set(valuePtr.Elem())

	return nil
}
