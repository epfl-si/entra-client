package client

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// GetToken get a token from Entra
func GetToken(clientID, clientSecret, tenantID string, restricted bool) (string, error) {

	var scope string
	if restricted {
		scope = "api://" + clientID + "/.default"
	} else {
		scope = "https://graph.microsoft.com/.default"
	}

	cred, err := azidentity.NewClientSecretCredential(
		tenantID,
		clientID,
		clientSecret,
		nil,
	)
	if err != nil {
		fmt.Printf("New Client Secret Credential Error: %s\n", err.Error())
		return "", err
	}

	result, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{
		Scopes: []string{
			scope,
		},
	})
	if err != nil {
		fmt.Printf("Get Token Error: %s\n", err.Error())
		return "", err
	}

	return result.Token, nil
}

// func GetClient(clientID, clientSecret, tenantID string) (*msgraphsdk.GraphServiceClient, error) {

// 	cred, _ := azidentity.NewClientSecretCredential(
// 		tenantID,
// 		clientID,
// 		clientSecret,
// 		nil,
// 	)
// 	fmt.Println("got cred")

// 	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
// 	if err != nil {
// 		fmt.Printf("getClient Error: %s\n", err.Error())
// 		return nil, err
// 	}
// 	fmt.Println("got client")

// 	return client, nil
// }
