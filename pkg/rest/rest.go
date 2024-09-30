// Package rest provides a REST client that can make requests to a REST API
package rest

import (
	"bytes"
	"net/http"
	"strings"
)

// Client is a REST client that can make requests to a REST API
type Client struct {
	// BaseURL is the base URL for the REST API
	BaseURL string
}

// Headers is a map of headers to include in the request
type Headers map[string]string

// New creates a new REST client
func New(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

// Get makes a GET request to the REST API
func (c *Client) Get(path string, headers Headers) (*http.Response, error) {
	var url string
	// if path does not start with BaseUrl, prepend it
	if path[0] != '/' {
		url = path
	} else {
		url = c.BaseURL + path
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	return http.DefaultClient.Do(req)
}

// Post makes a POST request to the REST API
func (c *Client) Post(path string, body []byte, headers Headers) (*http.Response, error) {
	url := c.BaseURL + path

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Delete makes a DELETE request to the REST API
func (c *Client) Delete(path string, headers Headers) (*http.Response, error) {
	url := c.BaseURL + path
	// create a new DELETE request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Patch makes a PATCH request to the REST API
func (c *Client) Patch(path string, body []byte, headers Headers) (*http.Response, error) {
	url := c.getNormalizedpath(path)
	// create a new PATCH request
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Put makes a PUT request to the REST API
func (c *Client) Put(path string, body []byte, headers Headers) (*http.Response, error) {
	url := c.BaseURL + path
	// create a new PUT request
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// TokenBearerString returns a Headerstring ct with the Authorization header set to the given token
func TokenBearerString(token string) string {
	return "Bearer " + token
}

func (c *Client) getNormalizedpath(path string) string {
	var url string
	// if path does not start with BaseUrl, prepend it
	if path[0] != '/' {
		url = path
	} else {
		url = c.BaseURL + path
	}

	return url
}

func refreshTokenNeeded(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "the token is expired") {
			// refresh token
			return true

		}

		return true
	}
	return false
}
