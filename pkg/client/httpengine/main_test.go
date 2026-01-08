package httpengine

import (
	"testing"

	"github.com/epfl-si/entra-client/pkg/client/models"
	"github.com/epfl-si/entra-client/pkg/rest"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		opts      []Option
		assertion assert.ErrorAssertionFunc
		validate  func(*testing.T, *HTTPClient)
	}{
		{
			name:      "creates client with default logger",
			opts:      nil,
			assertion: assert.NoError,
			validate: func(t *testing.T, c *HTTPClient) {
				assert.NotNil(t, c.Log, "logger should not be nil")
				assert.Equal(t, "https://graph.microsoft.com/v1.0", c.BaseURL)
				assert.NotNil(t, c.RestClient, "rest client should be initialized")
				assert.NotNil(t, c.appIDCache, "appID cache should be initialized")
				assert.NotNil(t, c.spIDCache, "spID cache should be initialized")
				assert.False(t, c.More, "More should be false by default")
			},
		},
		{
			name: "creates client with custom logger",
			opts: []Option{
				WithLogger(zap.NewNop()),
			},
			assertion: assert.NoError,
			validate: func(t *testing.T, c *HTTPClient) {
				assert.NotNil(t, c.Log, "logger should not be nil")
				// The logger should be the nop logger we passed
				assert.Equal(t, zap.NewNop(), c.Log)
				assert.Equal(t, "https://graph.microsoft.com/v1.0", c.BaseURL)
				assert.NotNil(t, c.RestClient, "rest client should be initialized")
			},
		},
		{
			name: "creates client with custom development logger",
			opts: func() []Option {
				logger := zap.Must(zap.NewDevelopment())
				return []Option{WithLogger(logger)}
			}(),
			assertion: assert.NoError,
			validate: func(t *testing.T, c *HTTPClient) {
				assert.NotNil(t, c.Log, "logger should not be nil")
				assert.NotNil(t, c.RestClient, "rest client should be initialized")
				assert.NotNil(t, c.appIDCache, "appID cache should be initialized")
				assert.NotNil(t, c.spIDCache, "spID cache should be initialized")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.opts...)
			tt.assertion(t, err)
			if err == nil && tt.validate != nil {
				tt.validate(t, got)
			}
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
