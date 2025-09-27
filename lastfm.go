// Package lastfmgo implements the [last.fm API]
//
// [last.fm API]: https://www.last.fm/api
package lastfmgo

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Lastfm struct {
	Key      string `json:"key"`
	Secret   string
	Username string
}

// SubmitLastfmCommand submits a request to the last.fm API
// the period is the period of data to return
// the apiKey is specific to your app, see the README for instructions
// the string returned is JSON
func SubmitLastfmCommand(period string, apiKey string, user string) (string, error) {
	apiURLBase := "https://ws.audioscrobbler.com/2.0/?"
	queryParameters := url.Values{}
	queryParameters.Set("method", "user.gettopartists")
	queryParameters.Set("user", user)
	switch period {
	case "weekly":
		queryParameters.Set("period", "7day")
	case "annual":
		queryParameters.Set("period", "12month")
	case "quarterly":
		queryParameters.Set("period", "3month")
	}
	queryParameters.Set("api_key", apiKey)
	queryParameters.Set("format", "json")
	fullURL := apiURLBase + queryParameters.Encode()
	lastfmResponse, statusCode, err := WebGet(fullURL)
	if err != nil {
		fmt.Println(statusCode)
		return lastfmResponse, err
	}
	return lastfmResponse, err
}

// webGet handles contacting a URL
func WebGet(url string) (string, int, error) {
	response, err := http.Get(url)
	if err != nil {
		return "Error accessing URL", 0, err
	}
	result, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		statusCodeString := fmt.Sprintf("Response failed with status code: %d and \nbody: %s\n", response.StatusCode, result)
		fmt.Println(statusCodeString)
		panic("Invalid status, data will be garbage")
	}
	if err != nil {
		return "Error reading response", 0, err
	}
	return string(result), response.StatusCode, err

}
