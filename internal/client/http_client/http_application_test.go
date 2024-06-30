package http_client_test

import (
	"epfl-entra/internal/client/http_client"
	"epfl-entra/internal/models"
	"epfl-entra/pkg/rest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHTTPClient_GetApplication(t *testing.T) {
	type args struct {
		id   string
		opts models.ClientOptions
	}
	tests := []struct {
		name string
		args args
		want *models.Application
		err  error
	}{
		{
			"Get Dev Portal RCP SAML22 application by ID",
			args{"4338fbfb-83b6-44be-ab56-7bb5e1f91b86", models.ClientOptions{Select: "id,displayName"}},
			&models.Application{ID: "4338fbfb-83b6-44be-ab56-7bb5e1f91b86", DisplayName: "Dev Portal RCP SAML2"},
			nil,
		},
		{
			"Non existant ID causes empty application",
			args{"NOWAYj", models.ClientOptions{Select: "id,displayName"}},
			&models.Application{ID: "", DisplayName: ""},
			nil,
		},
	}

	godotenv.Load("../../../.env")
	Client, err := http_client.New()
	if err != nil {
		t.Log("Testing require working environment variables")
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.GetApplication(tt.args.id, tt.args.opts)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestHTTPClient_GetApplications(t *testing.T) {
	type fields struct {
		AccessToken string
		BaseURL     string
		More        bool
		Secret      string
		ClientID    string
		Tenant      string
		RestClient  *rest.Client
		Log         *zap.Logger
	}
	type args struct {
		opts models.ClientOptions
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []*models.Application
		want1     string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}

	godotenv.Load("../../../.env")
	Client, err := http_client.New()
	if err != nil {
		t.Log("Testing require working environment variables")
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Client.GetApplications(tt.args.opts)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
