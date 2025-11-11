// Package models provides the models for the AppRole
// AppRole is part of an Application but assignable to Service Principals
// https://learn.microsoft.com/en-us/graph/api/resources/approle?view=graph-rest-1.0
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
