package models

// OptionalClaim holds optional claims definitions
type OptionalClaim struct {
	Name                 string   `json:"name,omitempty"`
	Essential            bool     `json:"essential,omitempty"`
	Source               string   `json:"source,omitempty"`
	AdditionalProperties []string `json:"additionalProperties,omitempty"`
}
