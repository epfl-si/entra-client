// Package models provides the models for the application
package models

// OAuth2PermissionGrant represents a delegated permission grant authorizing a client
// service principal to access an API on behalf of a user.
//
// Resources: https://learn.microsoft.com/en-us/graph/api/resources/oauth2permissiongrant?view=graph-rest-1.0
type OAuth2PermissionGrant struct {
	ID          string `json:"id,omitempty"`
	ClientID    string `json:"clientId,omitempty"`
	ConsentType string `json:"consentType,omitempty"`
	PrincipalID string `json:"principalId,omitempty"`
	ResourceID  string `json:"resourceId,omitempty"`
	Scope       string `json:"scope,omitempty"`
}

// OAuth2PermissionGrantResponse represents the response from the Graph API when listing grants.
type OAuth2PermissionGrantResponse struct {
	Context  string                  `json:"@odata.context"`
	NextLink string                  `json:"@odata.nextLink"`
	Value    []OAuth2PermissionGrant `json:"value"`
}
