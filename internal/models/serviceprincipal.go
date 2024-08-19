// Package models provides the models for the application
package models

import "time"

// Info represents the info of a service principal
type Info struct {
	LogoURL             string `json:"logoUrl,omitempty"`
	MarketingURL        string `json:"marketingUrl,omitempty"`
	PrivacyStatementURL string `json:"privacyStatementUrl,omitempty"`
	SupportURL          string `json:"supportUrl,omitempty"`
	TermsOfServiceURL   string `json:"termsOfServiceUrl,omitempty"`
}

// OAuth2PermissionScope represents an OAuth2 permission scope
type OAuth2PermissionScope struct {
	AdminConsentDescription string `json:"adminConsentDescription,omitempty"`
	AdminConsentDisplayName string `json:"adminConsentDisplayName,omitempty"`
	ID                      string `json:"id,omitempty"`
	IsEnabled               bool   `json:"isEnabled,omitempty"`
	Type                    string `json:"type,omitempty"`
	UserConsentDescription  string `json:"userConsentDescription,omitempty"`
	UserConsentDisplayName  string `json:"userConsentDisplayName,omitempty"`
	Value                   string `json:"value,omitempty"`
}

// VerifiedPublisher represents a verified publisher
type VerifiedPublisher struct {
	DisplayName         string     `json:"displayName,omitempty"`
	VerifiedPublisherID string     `json:"verifiedPublisherId,omitempty"`
	AddedDateTime       *time.Time `json:"addedDateTime,omitempty"`
}

// ServicePrincipal represents the part of an application that can be instanciated in several tenants
// (it relates to the Application->Enterprise applications menu in Entra)
//
// Resources: https://learn.microsoft.com/en-us/graph/api/resources/serviceprincipal?view=graph-rest-1.0
type ServicePrincipal struct {
	ID                                 string                  `json:"id,omitempty"`
	DeletedDateTime                    *time.Time              `json:"deletedDateTime,omitempty"`
	AccountEnabled                     bool                    `json:"accountEnabled,omitempty"`
	AppID                              string                  `json:"appId,omitempty"`
	ApplicationTemplateID              string                  `json:"applicationTemplateId,omitempty"`
	AppDisplayName                     string                  `json:"appDisplayName,omitempty"` // Read-only
	AlternativeNames                   []string                `json:"alternativeNames,omitempty"`
	AppOwnerOrganizationID             string                  `json:"appOwnerOrganizationId,omitempty"`
	DisplayName                        string                  `json:"displayName,omitempty"`
	AppRoleAssignmentRequired          bool                    `json:"appRoleAssignmentRequired,omitempty"`
	LoginURL                           string                  `json:"loginUrl,omitempty"`
	LogoutURL                          string                  `json:"logoutUrl,omitempty"`
	Homepage                           string                  `json:"homepage,omitempty"`
	NotificationEmailAddresses         []string                `json:"notificationEmailAddresses,omitempty"`
	PreferredSingleSignOnMode          string                  `json:"preferredSingleSignOnMode,omitempty"`
	PreferredTokenSigningKeyThumbprint string                  `json:"preferredTokenSigningKeyThumbprint,omitempty"`
	ReplyUrls                          []string                `json:"replyUrls,omitempty"`
	ServicePrincipalNames              []string                `json:"servicePrincipalNames,omitempty"`
	ServicePrincipalType               string                  `json:"servicePrincipalType,omitempty"`
	Tags                               []string                `json:"tags,omitempty"`
	TokenEncryptionKeyID               string                  `json:"tokenEncryptionKeyId,omitempty"`
	SamlSingleSignOnSettings           interface{}             `json:"samlSingleSignOnSettings,omitempty"`
	AddIns                             []interface{}           `json:"addIns,omitempty"`
	AppRoles                           []AppRole               `json:"appRoles,omitempty"`
	Info                               *Info                   `json:"info,omitempty"`
	KeyCredentials                     []KeyCredential         `json:"keyCredentials,omitempty"`
	OAuth2PermissionScopes             []OAuth2PermissionScope `json:"oauth2PermissionScopes,omitempty"`
	PasswordCredentials                []PasswordCredential    `json:"passwordCredentials,omitempty"`
	VerifiedPublisher                  *VerifiedPublisher      `json:"verifiedPublisher,omitempty"`
}
