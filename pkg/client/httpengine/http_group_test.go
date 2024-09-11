package httpengine

import (
	"testing"

	"github.com/epfl-si/entra-client/pkg/client/models"
	"github.com/epfl-si/entra-client/pkg/rest"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHTTPClient_GetGroup(t *testing.T) {
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
		want      *models.Group
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
			got, err := c.GetGroup(tt.args.id, tt.args.opts)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHTTPClient_GetGroups(t *testing.T) {
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
		want      []*models.Group
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
			got, got1, err := c.GetGroups(tt.args.opts)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
