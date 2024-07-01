package httpengine

import (
	"epfl-entra/internal/models"
	"epfl-entra/pkg/rest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		want      *HTTPClient
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHTTPClient_buildHeaders(t *testing.T) {
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
		name   string
		fields fields
		args   args
		want   rest.Headers
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
			assert.Equal(t, tt.want, c.buildHeaders(tt.args.opts))
		})
	}
}

func Test_buildQueryString(t *testing.T) {
	type args struct {
		opts models.ClientOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, buildQueryString(tt.args.opts))
		})
	}
}
