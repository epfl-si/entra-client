package models

// RequiredResource represents a required resource
type RequiredResource struct {
	ResourceAppID  string           `json:"resourceAppId,omitempty"`
	ResourceAccess []ResourceAccess `json:"resourceAccess,omitempty"`
}
