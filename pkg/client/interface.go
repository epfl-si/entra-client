// Package client provides the client interface for the application
package client

import (
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/crewjam/saml"
)

// Service is the interface for the client
type Service interface {
	// Utility
	CreateOIDCApplication(app *models.Application, appOptions *models.AppOptions) (newApp *models.Application, newSP *models.ServicePrincipal, secret string, err error)
	CreatePortalApplication(app *models.Application, options models.ClientOptions) (newApp *models.Application, newSP *models.ServicePrincipal, err error)
	GetClientID() (tenant string)
	GetSecret() (tenant string)
	GetTenant() (tenant string)
	GetToken(restricted bool) (token string, err error)

	// Application Template
	GetApplicationTemplate(id string, options models.ClientOptions) (apps *models.ApplicationTemplate, err error)
	GetApplicationTemplates(options models.ClientOptions) (apps []*models.ApplicationTemplate, nextURL string, err error)

	InstantiateApplicationTemplate(id, name string, options models.ClientOptions) (application *models.Application, servicePrincipal *models.ServicePrincipal, err error)

	// Application
	AddClaimToApplication(id, name, source, location string, basicPreset bool, options models.ClientOptions) (err error)
	AddPasswordToApplication(id, keyName string, options models.ClientOptions) (password *models.PasswordCredential, err error)
	CreateApplication(app *models.Application, options models.ClientOptions) (newApp *models.Application, err error)
	DeleteApplication(id string, options models.ClientOptions) (err error)
	GetApplication(id string, options models.ClientOptions) (apps *models.Application, err error)
	GetApplicationByAppID(id string, options models.ClientOptions) (apps *models.Application, err error)
	GetApplications(options models.ClientOptions) (apps []*models.Application, nextURL string, err error)
	GiveConsentToApplication(spObjectID string, scopes []string, options models.ClientOptions) (err error)
	PatchApplication(id string, app *models.Application, options models.ClientOptions) (err error)
	PatchApplicationTokenGroup(id string, groupClaimName string, opts models.ClientOptions) (err error)
	WaitApplication(id string, timeout int, options models.ClientOptions) (err error)
	//GiveConsentToApplication(id string, options models.ClientOptions) (err error)
	GetApplicationConsents(options models.ClientOptions) (body string, err error)

	// Claims Mapping Policy
	AssignClaimsMappingPolicy(cmpID, spID string, options models.ClientOptions) (err error)
	CreateClaimsMappingPolicy(app *models.ClaimsMappingPolicy, options models.ClientOptions) (id string, err error)
	DeleteClaimsMappingPolicy(id string, options models.ClientOptions) (err error)
	GetClaimsMappingPolicies(options models.ClientOptions) (groups []*models.ClaimsMappingPolicy, nextURL string, err error)
	GetClaimsMappingPolicy(cmpID string, options models.ClientOptions) (claimsMappingPolicy *models.ClaimsMappingPolicy, err error)
	GetClaimsMappingPolicyByAppID(appID string, options models.ClientOptions) (claimsMappingPolicy *models.ClaimsMappingPolicy, err error)
	ListUsageClaimsMappingPolicy(cmpID string, options models.ClientOptions) (groups []*models.DirectoryObject, err error)
	UnassignClaimsMappingPolicy(spID, cmpID string, options models.ClientOptions) (err error)
	PatchClaimsMappingPolicy(cmpID string, app *models.ClaimsMappingPolicy, options models.ClientOptions) (err error)

	// Extension
	GetExtension(id string, options models.ClientOptions) (extension *models.ExtensionProperty, err error)
	GetExtensions(options models.ClientOptions) (extensions []*models.ExtensionProperty, err error)

	// Group
	CreateGroup(app *models.Group, options models.ClientOptions) (err error)
	DeleteGroup(id string, options models.ClientOptions) (err error)
	GetGroup(id string, options models.ClientOptions) (groups *models.Group, err error)
	GetGroups(options models.ClientOptions) (groups []*models.Group, nextURL string, err error)

	// Service Principal
	AddCertificateToServicePrincipal(servicePrincipalID string, base64 string, options models.ClientOptions) (err error)
	//   Service Principal user/group management
	AddGroupToServicePrincipal(servicePrincipalID, groupID string, options models.ClientOptions) (err error)
	RemoveGroupFromServicePrincipal(servicePrincipalID, groupID string, options models.ClientOptions) (err error)
	GetGroupsFromServicePrincipal(servicePrincipalID string, options models.ClientOptions) (groups []*models.Group, err error)
	GetAssignmentsFromServicePrincipal(servicePrincipalI string, options models.ClientOptions) (assignment []*models.AppRoleAssignment, err error)
	AssignAppRoleToServicePrincipal(assignment *models.AppRoleAssignment, options models.ClientOptions) (err error)

	AddKeyToServicePrincipal(servicePrincipalID string, keyCredential saml.KeyDescriptor, options models.ClientOptions) (err error)

	AssignClaimsPolicyToServicePrincipal(claimsPolicyID, servicePrincipalID string) (err error)
	CreateServicePrincipal(app *models.ServicePrincipal, options models.ClientOptions) (newServicePrincipal *models.ServicePrincipal, err error)
	DeleteServicePrincipal(id string, options models.ClientOptions) (err error)

	GetClaimsMappingPoliciesForServicePrincipal(id string, options models.ClientOptions) (claimsMappingPolicies []*models.ClaimsMappingPolicy, nextURL string, err error)
	GetServicePrincipal(id string, options models.ClientOptions) (serviceprincipal *models.ServicePrincipal, err error)
	GetServicePrincipalByAppID(id string, options models.ClientOptions) (serviceprincipal *models.ServicePrincipal, err error)
	GetServicePrincipals(options models.ClientOptions) (serviceprincipals []*models.ServicePrincipal, nextURL string, err error)
	PatchServicePrincipal(id string, app *models.ServicePrincipal, options models.ClientOptions) (err error)
	UnassignClaimsPolicyFromServicePrincipal(claimsPolicyID, servicePrincipalID string, options models.ClientOptions) (err error)
	WaitServicePrincipal(id string, timeout int, options models.ClientOptions) (err error)

	// User
	CreateUser(app *models.User, options models.ClientOptions) (err error)
	DeleteUser(id string, options models.ClientOptions) (err error)
	GetUser(id string, options models.ClientOptions) (app *models.User, err error)
	GetUsers(options models.ClientOptions) (users []*models.User, nextURL string, err error)
	UpdateUser(app *models.User, options models.ClientOptions) (err error)

	// Identity
	AddApplicationToAuthenticationEventListeners(AuthenticationEventListenersId string, IncludeApplications *models.IdentityAuthenticationEventListenersIncludeApplicationsBody, opts models.ClientOptions) (err error)
	RemoveApplicationToAuthenticationEventListeners(AuthenticationEventListenersId string, appId string, opts models.ClientOptions) (err error)
}
