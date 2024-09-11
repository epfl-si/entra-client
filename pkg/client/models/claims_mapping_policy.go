package models

// ClaimsMappingPolicy represents the claims mapping policy
type ClaimsMappingPolicy struct {
	Definition            []string `json:"definition,omitempty"`
	DisplayName           string   `json:"displayName,omitempty"`
	ID                    string   `json:"id,omitempty"`
	IsOrganizationDefault bool     `json:"isOrganizationDefault,omitempty"`
}
