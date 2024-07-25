// Package models provides the models for the application
package models

type ApplicationTemplateResponse struct {
	Context  string                 `json:"@odata.context"`
	NextLink string                 `json:"@odata.nextLink"`
	Value    []*ApplicationTemplate `json:"value"`
}
