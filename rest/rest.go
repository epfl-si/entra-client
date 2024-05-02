package rest

import (
	"bytes"
	"net/http"
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
	url := c.BaseURL + path
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
func (c *Client) Post(path string, body []byte) (*http.Response, error) {
	url := c.BaseURL + path
	return http.Post(url, "application/json", bytes.NewReader(body))
}

// TokenHeaders returns a Headerstring ct with the Authorization header set to the given token
func TokenBearerString(token string) string {
	return "Bearer " + token
}
