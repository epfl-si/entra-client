// Package models provides the models for the application
package models

import "time"

// AppRoleAssignment is the struct used to assign user/group to application principal
//
// Resources: https://learn.microsoft.com/en-us/graph/api/resources/approleassignment?view=graph-rest-1.0
type AppRoleAssignment struct {
	ID                   string     `json:"id,omitempty"`
	DeletedDateTime      *time.Time `json:"deletedDateTime,omitempty"`
	AppRoleID            string     `json:"appRoleId,omitempty"`
	CreatedDateTime      *time.Time `json:"createdDateTime,omitempty"`
	PrincipalDisplayName string     `json:"principalDisplayName,omitempty"`
	PrincipalID          string     `json:"principalId,omitempty"`
	PrincipalType        string     `json:"principalType,omitempty"`
	ResourceDisplayName  string     `json:"resourceDisplayName,omitempty"`
	ResourceID           string     `json:"resourceId,omitempty"`
}

// AppRoleAssignmentResponse represents the response from the Graph API when listing app role assignments
type AppRoleAssignmentResponse struct {
	Context  string              `json:"@odata.context"`
	NextLink string              `json:"@odata.nextLink"`
	Value    []AppRoleAssignment `json:"value"`
}
