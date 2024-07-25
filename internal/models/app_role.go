// Package models provides the models for the application
package models

// AppRole represents the app role
type AppRole struct {
	AllowedMemberTypes []string `json:"allowedMemberTypes,omitempty"`
	Description        string   `json:"description,omitempty"`
	DisplayName        string   `json:"displayName,omitempty"`
	ID                 string   `json:"id,omitempty"`
	IsEnabled          bool     `json:"isEnabled,omitempty"`
	Origin             string   `json:"origin,omitempty"`
	// Value              *string  `json:"value"`
	Value string `json:"value"`
}
