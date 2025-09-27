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

// User API endpoints

// userGetTopArtists accesses the endpoint defined as user.getTopArtists
// user  : The user name to fetch top artists for.
// period : overall | 7day | 1month | 3month | 6month | 12month - The time period over which to retrieve top artists for.
// limit : The number of results to fetch per page. Defaults to 50.
// page : The page number to fetch. Defaults to first page.
// apiKey : A Last.fm API key. See README for instructions for obtaining one.
// returns the JSON response from the last.fm API and/or any golang errors
func UserGetTopArtists(user string, period string, limit string, page string, apiKey string) (string, error) {
	parameters := map[string]string{
		"user":    user,
		"period":  period,
		"limit":   limit,
		"page":    page,
		"api_key": apiKey,
	}
	return submitLastfmCommand(parameters)
}

// submitLastfmCommand submits a request to the last.fm API
// parameters is a map in which the key is the name of the parameter and the value is the value of the parameter
// the string returned is JSON
func submitLastfmCommand(parameters map[string]string) (string, error) {
	apiURLBase := "https://ws.audioscrobbler.com/2.0/?"
	queryParameters := url.Values{}
	for key, value := range parameters {
		queryParameters.Set(key, value)
	}
	queryParameters.Set("format", "json")
	fullURL := apiURLBase + queryParameters.Encode()
	lastfmResponse, statusCode, err := webGet(fullURL)
	if err != nil {
		fmt.Println(statusCode)
		return lastfmResponse, err
	}
	return lastfmResponse, err
}

// webGet handles contacting a URL
func webGet(url string) (string, int, error) {
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
