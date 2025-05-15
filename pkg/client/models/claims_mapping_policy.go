package models

// ClaimsMappingPolicy represents the claims mapping policy
type ClaimsMappingPolicy struct {
	Definition            []string `json:"definition,omitempty"`
	DisplayName           string   `json:"displayName,omitempty"`
	ID                    string   `json:"id,omitempty"`
	IsOrganizationDefault bool     `json:"isOrganizationDefault,omitempty"`
}

type ClaimsMappingPolicyEpfl struct {
	ID             string `json:"id" gorm:"column:id"`
	Base           bool   `json:"base" gorm:"column:base"`
	Cfs            bool   `json:"cfs" gorm:"column:cfs"`
	Authorizations bool   `json:"authorizations" gorm:"column:authorizations"`
	Accreds        bool   `json:"accreds" gorm:"column:accreds"`
}
