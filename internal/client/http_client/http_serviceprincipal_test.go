package httpengine

import (
	"epfl-entra/internal/models"
	"epfl-entra/pkg/rest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHTTPClient_AssociateAppRoleToServicePrincipal(t *testing.T) {
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
		assignment *models.AppRoleAssignment
		opts       models.ClientOptions
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
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
			tt.assertion(t, c.AssignAppRoleToServicePrincipal(tt.args.assignment, tt.args.opts))
		})
	}
}
