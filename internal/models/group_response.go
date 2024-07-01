// Package models provides the models for the application
package models

type GroupResponse struct {
	Context  string   `json:"@odata.context"`
	NextLink string   `json:"@odata.nextLink"`
	Value    []*Group `json:"value"`
}
