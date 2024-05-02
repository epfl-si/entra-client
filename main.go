package main

import (
	"context"
	"fmt"
	"os"

	"entra-client/rest"

	"github.com/joho/godotenv"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

func main() {

	godotenv.Load()

	secret := os.Getenv("ENTRA_SECRET")
	if secret == "" {
		panic("ENTRA_SECRET is not set")
	}

	tenant := os.Getenv("ENTRA_TENANT")
	if tenant == "" {
		panic("ENTRA_TENANT is not set")
	}

	clientID := os.Getenv("ENTRA_CLIENTID")
	if clientID == "" {
		panic("ENTRA_CLIENTID is not set")
	}

	accessToken := os.Getenv("ENTRA_ACCESS_TOKEN")
	if accessToken == "" {
		panic("ENTRA_ACCESS_TOKEN is not set")
	}

	// confidential clients have a credential, such as a secret or a certificate
	cred, err := confidential.NewCredFromSecret(secret)
	if err != nil {
		panic(err)
	}
	client, err := confidential.New("https://login.microsoftonline.com/"+tenant, clientID, cred)
	if err != nil {
		panic(err)
	}

	scopes := []string{"https://graph.microsoft.com/.default"}
	result, err := client.AcquireTokenSilent(context.TODO(), scopes)
	if err != nil {
		// cache miss, authenticate with another AcquireToken... method
		fmt.Println("Cache miss")
		result, err = client.AcquireTokenByCredential(context.TODO(), scopes)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("Access Token: %s\n", result.AccessToken)
	fmt.Printf("Using  Token: %s\n", accessToken)
	accessToken = result.AccessToken
	fmt.Printf("\ncurl \"https://graph.microsoft.com/v1.0/groups\" -H \"Authorization: %s\"", rest.TokenBearerString(accessToken))

	fmt.Println("Creating the client")
	r := rest.New("https://graph.microsoft.com/v1.0")
	fmt.Println("Calling the Microsoft Graph API...")
	response, err := r.Get("/groups", rest.Headers{"Authorization": rest.TokenBearerString(accessToken)})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
