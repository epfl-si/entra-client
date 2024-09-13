// Package client provides the client interface for the application
package client

import (
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/crewjam/saml"
)

type Service interface {
	GetToken() (token string)
	CreatePortalApplication(app *models.Application, options models.ClientOptions) (newApp *models.Application, newSP *models.ServicePrincipal, err error)
	CreateOIDCApplication(app *models.Application) (newApp *models.Application, newSP *models.ServicePrincipal, secret string, err error)


	InstantiateApplicationTemplate(id, name string, options models.ClientOptions) (application *models.Application, servicePrincipal *models.ServicePrincipal, err error)
	GetApplicationTemplate(id string, options models.ClientOptions) (apps *models.ApplicationTemplate, err error)
	GetApplicationTemplates(options models.ClientOptions) (apps []*models.ApplicationTemplate, nextURL string, err error)

	AddPasswordToApplication(id, keyName string, options models.ClientOptions) (password *models.PasswordCredential, err error)
	CreateApplication(app *models.Application, options models.ClientOptions) (newApp *models.Application, err error)
	DeleteApplication(id string, options models.ClientOptions) (err error)
	GetApplication(id string, options models.ClientOptions) (apps *models.Application, err error)
	GetApplications(options models.ClientOptions) (apps []*models.Application, nextURL string, err error)
	PatchApplication(id string, app *models.Application, options models.ClientOptions) (err error)
	WaitApplication(id string, timeout int, options models.ClientOptions) (err error)

	CreateClaimsMappingPolicy(app *models.ClaimsMappingPolicy, options models.ClientOptions) (id string, err error)
	DeleteClaimsMappingPolicy(id string, options models.ClientOptions) (err error)
	GetClaimsMappingPolicies(options models.ClientOptions) (groups []*models.ClaimsMappingPolicy, nextURL string, err error)

	CreateGroup(app *models.Group, options models.ClientOptions) (err error)
	DeleteGroup(id string, options models.ClientOptions) (err error)
	GetGroup(id string, options models.ClientOptions) (groups *models.Group, err error)
	GetGroups(options models.ClientOptions) (groups []*models.Group, nextURL string, err error)

	AddCertificateToServicePrincipal(servicePrincipalID string, base64 string, options models.ClientOptions) (err error)
	AddKeyToServicePrincipal(servicePrincipalID string, keyCredential saml.KeyDescriptor, options models.ClientOptions) (err error)
	AddGroupToServicePrincipal(servicePrincipalID, groupID string, options models.ClientOptions) (err error)
	AssignAppRoleToServicePrincipal(assignment *models.AppRoleAssignment, options models.ClientOptions) (err error)
	AssignClaimsPolicyToServicePrincipal(claimsPolicyID, servicePrincipalID string) (err error)
	CreateServicePrincipal(app *models.ServicePrincipal, options models.ClientOptions) (newServicePrincipal *models.ServicePrincipal, err error)
	PatchServicePrincipal(id string, app *models.ServicePrincipal, options models.ClientOptions) (err error)
	GetClaimsMappingPoliciesForServicePrincipal(id string, options models.ClientOptions) (claimsMappingPolicies []*models.ClaimsMappingPolicy, nextURL string, err error)
	GetServicePrincipal(id string, options models.ClientOptions) (serviceprincipal *models.ServicePrincipal, err error)
	GetServicePrincipals(options models.ClientOptions) (serviceprincipals []*models.ServicePrincipal, nextURL string, err error)
	GetAssignedAppRoles(id string, options models.ClientOptions) (assignment []*models.AppRoleAssignment, err error)
	UnassignClaimsPolicyFromServicePrincipal(claimsPolicyID, servicePrincipalID string, options models.ClientOptions) (err error)
	WaitServicePrincipal(id string, timeout int, options models.ClientOptions) (err error)

	CreateUser(app *models.User, options models.ClientOptions) (err error)
	DeleteUser(id string, options models.ClientOptions) (err error)
	GetUser(id string, options models.ClientOptions) (app *models.User, err error)
	GetUsers(options models.ClientOptions) (users []*models.User, nextURL string, err error)
	UpdateUser(app *models.User, options models.ClientOptions) (err error)
}
