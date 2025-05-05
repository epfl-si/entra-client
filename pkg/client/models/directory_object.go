package models

// DirectoryObject represents the Directory Object
type DirectoryObject struct {
	ID                     string   `json:"id,omitempty"`
	AccountEnabled         bool     `json:"accountEnabled,omitempty"`
	CreatedDateTime        string   `json:"createdDateTime,omitempty"`
	AppDisplayName         string   `json:"appDisplayName,omitempty"`
	AppID                  string   `json:"appId,omitempty"`
	AppOwnerOrganizationID string   `json:"appOwnerOrganizationId,omitempty"`
	PublisherName          string   `json:"publisherName,omitempty"`
	ServicePrincipalNames  []string `json:"servicePrincipalNames,omitempty"`
	ServicePrincipalType   string   `json:"servicePrincipalType,omitempty"`
	SignInAudience         string   `json:"signInAudience,omitempty"`
	DisplayName            string   `json:"displayName,omitempty"`
}
