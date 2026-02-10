package models

// ClaimsMappingPolicy represents the claims mapping policy
type ClaimsMappingPolicy struct {
	Definition            []string `json:"definition,omitempty"`
	DisplayName           string   `json:"displayName,omitempty"`
	ID                    string   `json:"id,omitempty"`
	IsOrganizationDefault bool     `json:"isOrganizationDefault,omitempty"`
}

type ClaimsMappingPolicyEpfl struct {
	ID      string `json:"id"`
	Base    bool   `json:"base"`
	Cfs     bool   `json:"cfs"`
	Accreds bool   `json:"accreds"`
}
