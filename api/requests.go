package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// REBRICKABLE_BASE_URL contains the Rebrickable API base URL
const REBRICKABLE_BASE_URL string = "https://rebrickable.com/api/v3/"

// createRequest returns a new http.Reqest object for the given url of the Rebrickable API
func createRequest(method string, url string, apiKey string, data url.Values) (*http.Request, error) {
	var body io.Reader
	if data != nil {
		body = strings.NewReader(data.Encode())
	}

	reqest, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	reqest.Header.Set("Accept", "application/json")
	reqest.Header.Set("Authorization", "key "+apiKey)

	return reqest, nil
}

// createGetRequest returns a new http.Reqest object for the given path of the Rebrickable API
func createGetRequest(url string, apiKey string) (*http.Request, error) {
	return createRequest(http.MethodGet, url, apiKey, nil)
}

// doRequest issues an HTTP request
func doRequest(client *http.Client, reqest *http.Request, v interface{}) error {
	log.Printf("Requesting %s\n", reqest.URL)

	var body []byte
	result, err := client.Do(reqest)
	if err != nil {
		return err
	}

	if result.Body != nil {
		defer result.Body.Close()
	}

	body, err = ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
