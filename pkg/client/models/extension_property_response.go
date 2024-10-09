// Package models provides the models for the application
package models

// ExtensionProperty represents an extension property
type ExtensionPropertyResponse struct {
	Context  string               `json:"@odata.context"`
	NextLink string               `json:"@odata.nextLink"`
	Value    []*ExtensionProperty `json:"value"`
}
