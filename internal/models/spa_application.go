package models

// SpaApplication represents the part of the OIDC configuration for SPA client
type SpaApplication struct {
	RedirectURIs []string `json:"redirectUris,omitempty"`
}
