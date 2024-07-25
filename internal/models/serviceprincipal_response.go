// Package models provides the models for the application
package models

type ServicePrincipalResponse struct {
	Context  string              `json:"@odata.context"`
	NextLink string              `json:"@odata.nextLink"`
	Value    []*ServicePrincipal `json:"value"`
}
