// Package httpengine provides engine using http client to interact with Microsoft Graph API
package httpengine

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	entraconfig "github.com/epfl-si/entra-client/internal/entra_config"
	client "github.com/epfl-si/entra-client/pkg/client"
	"github.com/epfl-si/entra-client/pkg/client/models"
	"github.com/epfl-si/entra-client/pkg/rest"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// HTTPClient is the client to the Microsoft Graph API
type HTTPClient struct {
	AccessToken string
	BaseURL     string
	More        bool
	Secret      string
	ClientID    string
	Tenant      string
	RestClient  *rest.Client
	Log         *zap.Logger
	EntraConfig *entraconfig.EntraConfig
}

// New creates a new HTTPClient
func New() (*HTTPClient, error) {
	var c HTTPClient

	c.More = false
	logger := zap.Must(zap.NewDevelopment())
	c.Log = logger
	c.BaseURL = "https://graph.microsoft.com/v1.0"
	c.RestClient = rest.New(c.BaseURL)

	err := c.GetConfig()
	if err != nil {
		return nil, err
	}

	if c.AccessToken == "" {
		// Get unrestricted token
		accessToken, err := client.GetToken(c.ClientID, c.Secret, c.Tenant, false)
		if err != nil {
			c.Log.Sugar().Debugf("New() - 0 - Error: %s\n", err.Error())
			return nil, err
		}
		c.AccessToken = accessToken
	}

	c.EntraConfig = entraconfig.New(c.GetTenant())

	return &c, nil
}

func (c *HTTPClient) buildHeaders(opts models.ClientOptions) rest.Headers {
	h := make(rest.Headers)

	h["Authorization"] = rest.TokenBearerString(c.AccessToken)

	if opts.Search != "" {
		h["ConsistencyLevel"] = "eventual"
	}

	return h
}

func buildQueryString(opts models.ClientOptions) string {
	qs := "?"

	var parameters []string

	if opts.Filter != "" {
		// uriencode opts.Filter
		opts.Filter = strings.ReplaceAll(opts.Filter, " ", "%20")
		parameters = append(parameters, "$filter="+opts.Filter)
		opts.Top = ""
	}
	if opts.Search != "" {
		parameters = append(parameters, "$search=\""+opts.Search+"\"")
	}

	if opts.Top != "" {
		parameters = append(parameters, "$top="+opts.Top)
	}

	if opts.Select != "" {
		parameters = append(parameters, "$select="+opts.Select)
	}

	if opts.Skip != "" {
		parameters = append(parameters, "$skip="+opts.Skip)
	}

	if opts.SkipToken != "" {
		parameters = append(parameters, "$skiptoken="+opts.SkipToken)
	}

	return qs + strings.Join(parameters, "&")
}

// GetConfig gets the configuration from the environment variables
func (c *HTTPClient) GetConfig() error {
	godotenv.Load()

	secret := os.Getenv("ENTRA_SECRET")
	if secret == "" {
		return errors.New("ENTRA_SECRET is not set")
	}
	c.Secret = secret

	tenant := os.Getenv("ENTRA_TENANT")
	if tenant == "" {
		return errors.New("ENTRA_TENANT is not set")
	}
	c.Tenant = tenant

	clientID := os.Getenv("ENTRA_CLIENTID")
	if clientID == "" {
		return errors.New("ENTRA_CLIENTID is not set")
	}
	c.ClientID = clientID

	// Using token from https://developer.microsoft.com/en-us/graph/graph-explorer
	// Accept empty token (will be retried by credentials)
	accessToken := os.Getenv("ENTRA_ACCESS_TOKEN")
	if accessToken != "" {
		c.AccessToken = accessToken
	}

	return nil
}

func getBody(response *http.Response) string {
	body, _ := io.ReadAll(io.Reader(response.Body))

	return string(body)
}

// GetToken returns the access token
func (c *HTTPClient) GetToken(restricted bool) (string, error) {
	return client.GetToken(c.ClientID, c.Secret, c.Tenant, restricted)
}

// GetCurrentToken returns the currently used access token
func (c *HTTPClient) GetCurrentToken() string {

	return c.AccessToken
}

// GetTenant returns the tenant
func (c *HTTPClient) GetTenant() string {

	return c.Tenant
}

// GetSecret returns the client secret
func (c *HTTPClient) GetSecret() string {

	return c.Secret
}

// GetClientID returns the client ID
func (c *HTTPClient) GetClientID() string {

	return c.ClientID
}

func normalizeThumbprint(thumbprint string) string {
	re, _ := regexp.Compile(`[\s\-]`)
	thumbprint = re.ReplaceAllString(thumbprint, "")

	return thumbprint
}

// CreatePortalApplication create an application and service principal
func (c *HTTPClient) CreatePortalApplication(app *models.Application, clientOptions models.ClientOptions) (*models.Application, *models.ServicePrincipal, error) {

	newApp, err := c.CreateApplication(app, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("CreateApplication: %w", err)
	}

	err = c.WaitApplication(newApp.ID, 60, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("WaitApplication: %w", err)
	}

	sp, err := c.CreateServicePrincipal(&models.ServicePrincipal{
		AppID:                newApp.AppID,
		ServicePrincipalType: "Application"}, clientOptions)

	if err != nil {
		return nil, nil, fmt.Errorf("CreateServicePrincipal: %W", err)
	}

	return newApp, sp, nil
}

// CreateOIDCApplication create an application and service principal
// redirect URIs if present must be
//
//	either in Web.RedirectURIs
//	or     in Spa.RedirectURIs
func (c *HTTPClient) CreateOIDCApplication(app *models.Application, appOptions *models.AppOptions) (newApp *models.Application, newSP *models.ServicePrincipal, secret string, err error) {
	var errs = ""
	// URIList := []models.URI{}

	bootstrApp := &models.Application{
		DisplayName: app.DisplayName,
	}

	if app.Web != nil && app.Web.RedirectURIs != nil {
		// Web application
		bootstrApp.Web = &models.WebSection{}
		bootstrApp.Web.RedirectURIs = app.Web.RedirectURIs
		// for i, uri := range app.Web.RedirectURIs {
		// 	URIList = append(URIList, models.URI{URI: uri, Index: i})
		// }
		// bootstrApp.Web.RedirectURISettings = URIList
	} else if app.Spa != nil && app.Spa.RedirectURIs != nil {
		// SPA application
		bootstrApp.Spa = &models.SpaApplication{
			RedirectURIs: app.Spa.RedirectURIs,
		}
	}

	opts := models.ClientOptions{}
	app, sp, err := c.CreatePortalApplication(bootstrApp, opts)
	if err != nil {
		return app, sp, "", err
	}

	scrt, err := c.AddPasswordToApplication(app.ID, app.DisplayName+"_secret", opts)
	if err != nil {
		errs += fmt.Sprintf("AddPasswordToApplication: %s\n", err.Error())
	}

	appPatch := &models.Application{
		RequiredResourceAccess: []models.RequiredResource{
			{
				ResourceAppID: c.EntraConfig.Get("MSGRAPH_API_RESOURCE_APP_ID"),
				ResourceAccess: []models.ResourceAccess{
					{
						ID:   c.EntraConfig.Get("MSGRAPH_EMAIL_RESOURCE_ID"),
						Type: "Scope",
					},
					{
						ID:   c.EntraConfig.Get("MSGRAPH_OFFLINE_ACCESS_RESOURCE_ID"),
						Type: "Scope",
					},
					{
						ID:   c.EntraConfig.Get("MSGRAPH_OPENID_RESOURCE_ID"),
						Type: "Scope",
					},
					{
						ID:   c.EntraConfig.Get("MSGRAPH_PROFILE_RESOURCE_ID"),
						Type: "Scope",
					},
					{
						ID:   c.EntraConfig.Get("MSGRAPH_USER_READ_RESOURCE_ID"),
						Type: "Scope",
					},
				},
			},
		},
	}
	appPatch.Web = &models.WebSection{
		ImplicitGrantSettings: &models.Grant{
			EnableIDTokenIssuance:     true,
			EnableAccessTokenIssuance: true,
		},
	}

	notes := "Unit: \nTest URL: https://login.microsoftonline.com/" + c.Tenant + "/oauth2/v2.0/authorize?client_id=" + app.AppID + "&response_type=token&redirect_uri=https://jwt.ms&scope=openid%20profile&state=12345&nonce=12345"
	appPatch.Notes = &notes
	version := 2
	t := true
	appPatch.API = &models.APIApplication{
		AcceptMappedClaims:          &t,
		RequestedAccessTokenVersion: &version,
	}

	// Causes error:
	// appPatch.AllowPublicClient = true
	appPatch.IsFallbackPublicClient = &t // For PKCE

	err = c.PatchApplication(app.ID, appPatch, opts)
	if err != nil {
		errs += fmt.Sprintf("PatchApplication: %s\n", err.Error())
	}

	//Waiting for the consent on DelegatedPermissionGrant.ReadWrite.All
	// Give consent to the application

	err = c.GiveConsentToApplication(sp.ID, []string{
		"User.Read",
		"openid",
		"profile",
		"email",
		"offline_access",
	}, opts)
	if err != nil {
		errs += fmt.Sprintf("GiveConsentToApplication: %s\n", err.Error())
	}

	// Configure claims (5th parameter is to add default claims)
	//err = c.AddClaimToApplication(app.ID, "", "", "", true, opts)

	// Customize application
	spPatch := &models.ServicePrincipal{}
	// setting Homepage default "Visible to all users" to true and is used for IdP initiated flows
	// sp.Homepage = "https://www.epfl.ch"
	spPatch.Tags = []string{"HideApp"} // If missing "Visible to all users" is true
	spPatch.AppRoleAssignmentRequired = true

	err = c.PatchServicePrincipal(sp.ID, spPatch, opts)
	if err != nil {
		errs += fmt.Sprintf("PatchServicePrincipal: %s\n", err.Error())
	}

	authorized := []string{}
	if appOptions == nil || appOptions.AuthorizedUsers == nil || len(appOptions.AuthorizedUsers) == 0 {
		authorized = []string{
			"AAD_All Hosts Users",
			"AAD_All Outside EPFL Users",
			"AAD_All Staff Users",
			"AAD_All Student Users",
		}
	} else {
		authorized = appOptions.AuthorizedUsers
	}

	for _, groupID := range authorized {
		err = c.AddGroupToServicePrincipal(sp.ID, groupID, opts)
		if err != nil {
			errs += fmt.Sprintf("AddGroupToServicePrincipal: %s\n", err.Error())
		}
	}

	// get default claims mapping policy
	cmps, _, err := c.GetClaimsMappingPolicies(models.ClientOptions{Filter: "isOrganizationDefault eq true"})
	if err != nil {
		errs += fmt.Sprintf("GetClaimsMappingPolicy: %s", err)
		c.Log.Sugar().Debugf("CreateOIDCApplication() - 0 - Error: %s\n", err.Error())
	}

	// If default claims mapping policy is found, assign it to the service principal
	if err == nil && len(cmps) == 1 {
		// assign default claims mapping policy to service principal
		err = c.AssignClaimsMappingPolicy(cmps[0].ID, sp.ID, opts)
		if err != nil {
			errs += fmt.Sprintf("AssignClaimsMappingPolicy: %s", err)
			c.Log.Sugar().Debugf("CreateOIDCApplication() - 1 - Error: %s\n", err.Error())
		}
	}

	if errs != "" {
		return app, sp, *scrt.SecretText, errors.New(errs)
	}

	return app, sp, *scrt.SecretText, nil

}
