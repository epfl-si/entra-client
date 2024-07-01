// Package models provides the models for the application
package models

type ApplicationResponse struct {
	Context  string         `json:"@odata.context"`
	NextLink string         `json:"@odata.nextLink"`
	Value    []*Application `json:"value"`
}
