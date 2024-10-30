// Package models provides the models for the application
package models

// ExtensionPropertyResponse represents an extension property reponse
type ExtensionPropertyResponse struct {
	Context  string               `json:"@odata.context"`
	NextLink string               `json:"@odata.nextLink"`
	Value    []*ExtensionProperty `json:"value"`
}
