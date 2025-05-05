// Package models provides the models for the application
package models

// ClaimsMappingPolicyListUsageResponse represents the response for listing claims mapping policy usage
type ClaimsMappingPolicyListUsageResponse struct {
	Context  string             `json:"@odata.context"`
	NextLink string             `json:"@odata.nextLink"`
	Value    []*DirectoryObject `json:"value"`
}
