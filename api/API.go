package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// REBRICKABLE_BASE_URL contains the Rebrickable API base URL
const REBRICKABLE_BASE_URL string = "https://rebrickable.com/api/v3/"

// I_API provides requestPage REST call for accessing Rebrickable's API
type I_API interface {
	requestPage(string, interface{}) error
}

// AbstractAPI is an abstract class for common API attributes
type AbstractAPI struct {
	client  *http.Client
	apiKey  string
	verbose bool
}

func (a *AbstractAPI) requestPage(url string, v interface{}) error {
	reqest, err := a.createGetRequest(url)
	if err != nil {
		return err
	}

	err = a.doRequest(reqest, v)
	if err != nil {
		return err
	}

	return nil
}

// createRequest returns a new http.Reqest object for the given url of the Rebrickable API
func (a *AbstractAPI) createRequest(method string, url string, data url.Values) (*http.Request, error) {
	var body io.Reader
	if data != nil {
		body = strings.NewReader(data.Encode())
	}

	reqest, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	reqest.Header.Set("Accept", "application/json")
	reqest.Header.Set("Authorization", "key "+a.apiKey)

	return reqest, nil
}

// createGetRequest returns a new http.Reqest object for the given path of the Rebrickable API
func (a *AbstractAPI) createGetRequest(url string) (*http.Request, error) {
	return a.createRequest(http.MethodGet, url, nil)
}

// doRequest issues an HTTP request
func (a *AbstractAPI) doRequest(reqest *http.Request, v interface{}) error {
	if a.verbose {
		log.Printf("Requesting %s", reqest.URL)
	}

	var body []byte
	result, err := a.client.Do(reqest)
	if err != nil {
		return err
	}
	if result.StatusCode != 200 {
		return fmt.Errorf("request returned '%s'", result.Status)
	}

	if result.Body != nil {
		defer result.Body.Close()
	}

	body, err = io.ReadAll(result.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
