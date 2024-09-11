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
