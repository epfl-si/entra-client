package models

// AppOptions are use to create an application
type AppOptions struct {
	DisplayName     string   `json:"displayName,omitempty"`
	LogoutURI       string   `json:"logoutUri,omitempty"`
	MetadataFile    string   `json:"metadataFile,omitempty"`
	RedirectURI     string   `json:"redirectUri,omitempty"`
	SAMLID          string   `json:"samlId,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	AuthorizedUsers []string `json:"authorizedUsers,omitempty"`
}
