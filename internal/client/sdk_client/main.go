package sdk_client

import (
	"context"

	"errors"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraph "github.com/microsoftgraph/msgraph-sdk-go"
	"go.uber.org/zap"
)

type SDKClient struct {
	AccessToken string
	Secret      string
	ClientID    string
	Tenant      string
	APIClient   *msgraph.GraphServiceClient
	Log         *zap.Logger
}

func (c *SDKClient) GetConfig() error {
	c.Secret = os.Getenv("ENTRA_SECRET")
	if c.Secret == "" {
		return errors.New("ENTRA_SECRET is not set")
	}

	c.Tenant = os.Getenv("ENTRA_TENANT")
	if c.Tenant == "" {
		return errors.New("ENTRA_TENANT is not set")
	}

	c.ClientID = os.Getenv("ENTRA_CLIENTID")
	if c.ClientID == "" {
		return errors.New("ENTRA_CLIENTID is not set")
	}

	c.AccessToken = os.Getenv("ENTRA_ACCESS_TOKEN")
	if c.AccessToken == "" {
		return errors.New("ENTRA_ACCESS_TOKEN is not set")
	}

	return nil
}

func New() (*SDKClient, error) {
	var c SDKClient

	c.GetConfig()
	c.AccessToken = ""
	if c.AccessToken == "" {
		accessToken, err := GetToken(c.ClientID, c.Secret, c.Tenant)
		if err != nil {
			return nil, err
		}
		c.AccessToken = accessToken
	}
	APIClient, err := GetClient(c.ClientID, c.Secret, c.Tenant)
	if err != nil {
		return nil, err
	}

	c.APIClient = APIClient

	return &c, nil
}

func getClient(clientID, clientSecret, tenantID string) (*msgraph.GraphServiceClient, error) {

	cred, _ := azidentity.NewClientSecretCredential(
		tenantID,
		clientID,
		clientSecret,
		nil,
	)
	fmt.Println("got cred")

	client, err := msgraph.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		fmt.Printf("getClient Error: %s\n", err.Error())
		return nil, err
	}
	fmt.Println("got client")

	return client, nil
}

func GetClient(clientID, clientSecret, tenantID string) (*msgraph.GraphServiceClient, error) {

	cred, _ := azidentity.NewClientSecretCredential(
		tenantID,
		clientID,
		clientSecret,
		nil,
	)

	client, err := msgraph.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetToken(clientID, clientSecret, tenantID string) (string, error) {

	cred, _ := azidentity.NewClientSecretCredential(
		tenantID,
		clientID,
		clientSecret,
		nil,
	)

	result, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{
		Scopes: []string{
			"https://graph.microsoft.com/.default",
		},
	})
	if err != nil {
		return "", err
	}

	return result.Token, nil
}

func safeString(s *string) string {
	fmt.Printf("safeString() - 0 - s: %s\n", *s)
	if s == nil {
		return ""
	}
	return *s
}
