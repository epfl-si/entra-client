package client

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

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
		fmt.Printf("Get Token Error: %s\n", err.Error())
		return "", err
	}

	// client, err := graph.NewGraphServiceClientWithCredentials(
	// 	cred, []string{"https://graph.microsoft.com/.default"})
	// if err != nil {
	// 	fmt.Printf("GetClient Error: %s\n", err.Error())
	// 	return "", err

	// }

	/// confidential clients have a credential, such as a secret or a certificate
	/*
		cred, err := confidential.NewCredFromSecret(secret)
		if err != nil {
			fmt.Printf("GetToken Error: %s\n", err.Error())
			return "", err
		}

		scopes := []string{"https://graph.microsoft.com/.default"}

		client, err := confidential.New("https://login.microsoftonline.com/"+tenant, clientID, cred)
		if err != nil {
			fmt.Printf("GetToken Error: %s\n", err.Error())
			return "", err

		}
	*/

	// cache miss, authenticate with another AcquireToken... method
	// fmt.Println("Cache miss")
	// result, err := client.AcquireTokenByCredential(context.TODO(), scopes)
	// if err != nil {
	// 	fmt.Printf("GetToken Error: %s\n", err.Error())
	// 	return "", err
	// }

	// fmt.Printf("GetToken: %+v\n", result)
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
