// Package models provides the models for the application
package models

import "time"

// ServicePrincipal represents the part of an application that can be instanciated in several tenants
// (it relates to the Application->Enterprise applications menu in Entra)
//
// Resources: https://learn.microsoft.com/en-us/graph/api/resources/serviceprincipal?view=graph-rest-1.0
type ServicePrincipal struct {
	AccountEnabled                     bool                    `json:"accountEnabled,omitempty"`
	AddIns                             []interface{}           `json:"addIns,omitempty"`
	AlternativeNames                   []string                `json:"alternativeNames,omitempty"`
	AppDescription                     string                  `json:"appDescription,omitempty"`
	AppDisplayName                     string                  `json:"appDisplayName,omitempty"` // Read-only
	AppID                              string                  `json:"appId,omitempty"`
	AppOwnerOrganizationID             string                  `json:"appOwnerOrganizationId,omitempty"`
	AppRoleAssignmentRequired          bool                    `json:"appRoleAssignmentRequired,omitempty"`
	AppRoles                           []AppRole               `json:"appRoles,omitempty"`
	ApplicationTemplateID              string                  `json:"applicationTemplateId,omitempty"`
	CustomSecurityAttributes           *string                 `json:"customSecurityAttributes,omitempty"`
	DeletedDateTime                    *time.Time              `json:"deletedDateTime,omitempty"`
	Description                        string                  `json:"description,omitempty"`
	DisplayName                        string                  `json:"displayName,omitempty"`
	DisabledByMicrosoftStatus          string                  `json:"disabledByMicrosoftStatus,omitempty"`
	Homepage                           string                  `json:"homepage,omitempty"`
	ID                                 string                  `json:"id,omitempty"`
	Info                               *Info                   `json:"info,omitempty"`
	KeyCredentials                     []KeyCredential         `json:"keyCredentials,omitempty"`
	LoginURL                           string                  `json:"loginUrl,omitempty"`
	LogoutURL                          string                  `json:"logoutUrl,omitempty"`
	Notes                              string                  `json:"notes,omitempty"`
	NotificationEmailAddresses         []string                `json:"notificationEmailAddresses,omitempty"`
	OAuth2PermissionScopes             []OAuth2PermissionScope `json:"oauth2PermissionScopes,omitempty"`
	PasswordCredentials                []PasswordCredential    `json:"passwordCredentials,omitempty"`
	PreferredSingleSignOnMode          string                  `json:"preferredSingleSignOnMode,omitempty"`
	PreferredTokenSigningKeyThumbprint string                  `json:"preferredTokenSigningKeyThumbprint,omitempty"`
	ReplyUrls                          []string                `json:"replyUrls,omitempty"`
	SamlSingleSignOnSettings           SAMLSingleSignOnSetting `json:"samlSingleSignOnSettings,omitempty"`
	ServicePrincipalNames              []string                `json:"servicePrincipalNames,omitempty"`
	ServicePrincipalType               string                  `json:"servicePrincipalType,omitempty"`
	SignInAudience                     string                  `json:"signInAudience,omitempty"`
	Tags                               []string                `json:"tags,omitempty"`
	TokenEncryptionKeyID               string                  `json:"tokenEncryptionKeyId,omitempty"`
	VerifiedPublisher                  *VerifiedPublisher      `json:"verifiedPublisher,omitempty"`
}
