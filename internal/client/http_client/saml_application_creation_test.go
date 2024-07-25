package httpengine

import (
	httpengine "epfl-entra/internal/client/http_client"
	"epfl-entra/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstantiateApplicationFromTempla(t *testing.T) {
	// This is a placeholder and should be replaced with actual setup, execution and assertion
	// ./ecli apptemplate instantiate --id 229946b9-a9fb-45b8-9531-efa47453ac9e --displayname "Test provisioning SAML AA"
	// ./ecli serviceprincipal patch --id b23f000f-8131-41ae-98b9-4279a8245ac2 --post '{"preferredSingleSignOnMode": "saml" }'
	// ./ecli application patch --id 3e3f000f-8131-41ae-98b9-4279a8245ac2 --post '{ "identifierUris": [ "https://signin.aws.amazon.com/saml" ], "web": { "redirectUris": [ "https://signin.aws.amazon.com/saml" ] }}'
	t.Run("test case 1", func(t *testing.T) {
		client, err := httpengine.New()
		if err != nil {
			t.Fatalf("error creating http client: %v", err)
		}
		opts := models.ClientOptions{}
		app, sp, err := client.InstantiateApplicationTemplate("229946b9-a9fb-45b8-9531-efa47453ac9e", "Automated test provisioning SAML AA", opts)
		assert.NoError(t, err)
		assert.NotNil(t, sp)
		assert.NotEmpty(t, sp.ID)
		err = client.WaitServicePrincipal(sp.ID, 60, opts)
		if err != nil {
			t.Fatalf("error waiting for application: %v", err)
		}
		err = client.PatchServicePrincipal(sp.ID, &models.ServicePrincipal{PreferredSingleSignOnMode: "saml"}, opts)
		assert.NoError(t, err)
		err = client.WaitApplication(app.ID, 60, opts)
		if err != nil {
			t.Fatalf("error waiting for application: %v", err)
		}
		err = client.PatchApplication(app.ID, &models.Application{IdentifierUris: []string{"https://signin.aws.amazon.com/saml"}, Web: &models.WebSection{RedirectURIs: []string{"https://signin.aws.amazon.com/saml"}}}, opts)
		assert.NoError(t, err)

	})
}
