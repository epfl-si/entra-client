// Package models provides the models for the application
package models

type ClaimsMappingPolicyResponse struct {
	Context  string                 `json:"@odata.context"`
	NextLink string                 `json:"@odata.nextLink"`
	Value    []*ClaimsMappingPolicy `json:"value"`
}
