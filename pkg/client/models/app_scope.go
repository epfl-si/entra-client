// Package models provides the models for the application
package models

// Scope is the struct used to get scopes
type ScopeDescription struct {
	AppID                  string                  `json:"appId,omitempty"`
	DisplayName            string                  `json:"displayName,omitempty"`
	Oauth2PermissionScopes []Oauth2PermissionScope `json:"oauth2PermissionScopes,omitempty"`
}

type Oauth2PermissionScope struct {
	ID                      string `json:"id,omitempty"`
	AdminConsentDescription string `json:"adminConsentDescription,omitempty"`
	AdminConsentDisplayName string `json:"adminConsentDisplayName,omitempty"`
	IsEnabled               bool   `json:"isEnabled,omitempty"`
	Type                    string `json:"type,omitempty"`
	UserConsentDescription  string `json:"userConsentDescription,omitempty"`
	UserConsentDisplayName  string `json:"userConsentDisplayName,omitempty"`
	Value                   string `json:"value,omitempty"`
}
