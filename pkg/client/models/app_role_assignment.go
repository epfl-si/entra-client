// Package models provides the models for the application
package models

// AppRoleAssignment is the struct used to assign user to application principal
type AppRoleAssignment struct {
	ID            string `json:"id"`
	PrincipalID   string `json:"principalId"`
	PrincipalType string `json:"principalType"`
	AppRoleID     string `json:"appRoleId"`
	ResourceID    string `json:"resourceId"`
}
