package api

import "net/http"

// I_API provides requestPage REST call for accessing Rebrickable's API
type I_API interface {
	requestPage(string, interface{}) error
}

// AbstractAPI is an abstract class for common API attributes
type AbstractAPI struct {
	client *http.Client
	apiKey string
}

func (a *AbstractAPI) requestPage(url string, v interface{}) error {
	reqest, err := CreateGetRequest(url, a.apiKey)
	if err != nil {
		return err
	}

	err = DoRequest(a.client, reqest, v)
	if err != nil {
		return err
	}

	return nil
}
