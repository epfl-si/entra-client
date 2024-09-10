package httpengine

import (
	"epfl-entra/pkg/entra-client/models"
	"epfl-entra/pkg/rest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHTTPClient_GetUser(t *testing.T) {
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
		id   string
		opts models.ClientOptions
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *models.User
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTTPClient{
				AccessToken: tt.fields.AccessToken,
				BaseURL:     tt.fields.BaseURL,
				More:        tt.fields.More,
				Secret:      tt.fields.Secret,
				ClientID:    tt.fields.ClientID,
				Tenant:      tt.fields.Tenant,
				RestClient:  tt.fields.RestClient,
				Log:         tt.fields.Log,
			}
			got, err := c.GetUser(tt.args.id, tt.args.opts)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHTTPClient_GetUsers(t *testing.T) {
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
		want      []*models.User
		want1     string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTTPClient{
				AccessToken: tt.fields.AccessToken,
				BaseURL:     tt.fields.BaseURL,
				More:        tt.fields.More,
				Secret:      tt.fields.Secret,
				ClientID:    tt.fields.ClientID,
				Tenant:      tt.fields.Tenant,
				RestClient:  tt.fields.RestClient,
				Log:         tt.fields.Log,
			}
			got, got1, err := c.GetUsers(tt.args.opts)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
