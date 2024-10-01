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
		accessToken, err := client.GetToken(c.ClientID, c.Secret, c.Tenant)
		if err != nil {
			c.Log.Sugar().Debugf("New() - 0 - Error: %s\n", err.Error())
		}
		c.AccessToken = accessToken
	}

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
func (c *HTTPClient) GetToken() string {

	return c.AccessToken
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
func (c *HTTPClient) CreateOIDCApplication(app *models.Application) (newApp *models.Application, newSP *models.ServicePrincipal, secret string, err error) {

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

	scrt, err := c.AddPasswordToApplication(app.ID, app.DisplayName+" secret", opts)
	if err != nil {
		return app, sp, "", err
	}

	appPatch := &models.Application{
		RequiredResourceAccess: []models.RequiredResource{
			{
				ResourceAppID: "00000003-0000-0000-c000-000000000000",
				ResourceAccess: []models.ResourceAccess{
					{
						ID:   "64a6cdd6-aab1-4aaf-94b8-3cc8405e90d0",
						Type: "Scope",
					},
					{
						ID:   "7427e0e9-2fba-42fe-b0c0-848c9e6a8182",
						Type: "Scope",
					},
					{
						ID:   "37f7f235-527c-4136-accd-4a02d197296e",
						Type: "Scope",
					},
					{
						ID:   "14dad69e-099b-42c9-810b-d002981feec1",
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

	version := 2
	t := true
	appPatch.API = &models.ApiApplication{
		AcceptMappedClaims:          &t,
		RequestedAccessTokenVersion: &version,
	}

	// Causes error:
	// appPatch.AllowPublicClient = true
	appPatch.IsFallbackPublicClient = &t // For PKCE

	err = c.PatchApplication(app.ID, appPatch, opts)
	if err != nil {
		return app, sp, "", err
	}

	// Configure claims (5th parameter is to add default claims)ß
	err = c.AddClaimToApplication(app.ID, "", "", "", true, opts)

	// Customize application
	spPatch := &models.ServicePrincipal{}
	// setting Homepage default "Visible to all users" to true and is used for IdP initiated flows
	// sp.Homepage = "https://www.epfl.ch"
	spPatch.Tags = []string{"HideApp"} // If missing "Visible to all users" is true
	spPatch.AppRoleAssignmentRequired = true

	err = c.PatchServicePrincipal(sp.ID, spPatch, opts)
	if err != nil {
		return app, sp, "", err
	}

	for _, groupID := range []string{
		"AAD_All Hosts Users",
		"AAD_All Outside EPFL Users",
		"AAD_All Staff Users",
		"AAD_All Student Users",
	} {

		err = c.AddGroupToServicePrincipal(sp.ID, groupID, opts)
		if err != nil {
			return app, sp, "", err
		}
	}

	// Works but can't be edited by portal
	// err = rootcmd.Client.AssignClaimsPolicyToServicePrincipal("b0a98d4a-221f-4d76-b6fb-7f6f0089175f", sp.ID)
	// if err != nil {
	// 	rootcmd.PrintErr(fmt.Errorf("Assign ClaimsPolicy %s to ServicePrincipal %s: %w", "b0a98d4a-221f-4d76-b6fb-7f6f0089175", sp.ID, err))
	// 	return
	// }

	return app, sp, *scrt.SecretText, nil

}
