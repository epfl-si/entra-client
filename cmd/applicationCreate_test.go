package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applicationCreate(t *testing.T) {

	// Transform the function into a test table
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
	}{
		{
			name:        "displayname is required",
			args:        []string{"application", "create"},
			expectedOut: "",
			expectedErr: "Name is required (use --displayname)\n",
		},
		{
			name:        "redirect_uri is required when displayname is provided",
			args:        []string{"application", "create", "--displayname", "test"},
			expectedOut: "",
			expectedErr: "Callback URL is required (use --redirect_uri)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rout, wout, oldout, rerr, werr, olderr := captureOutput()

			rootCmd.SetArgs(tt.args)
			rootCmd.Execute()

			out, err := releaseOutput(rout, wout, oldout, rerr, werr, olderr)

			assert.Equal(t, tt.expectedOut, string(out))
			assert.Equal(t, tt.expectedErr, string(err))
		})
	}

}
