package models

type APIApplication struct {
	AcceptMappedClaims          *bool                      `json:"acceptMappedClaims,omitempty"`
	KnownClientApplications     []string                   `json:"knownClientApplications,omitempty"`
	OAuth2PermissionScopes      []OAuth2PermissionScope    `json:"oauth2PermissionScopes,omitempty"`
	PreAuthorizedApplications   []PreAuthorizedApplication `json:"preAuthorizedApplications,omitempty"`
	RequestedAccessTokenVersion *int                       `json:"requestedAccessTokenVersion,omitempty"`
}
