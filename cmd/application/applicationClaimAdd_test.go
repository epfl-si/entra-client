package cmdapplication

import (
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_applicationClaimAdd(t *testing.T) {

	// Define test cases
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
	}{
		{
			name:        "ID is required",
			args:        []string{"application", "claim", "add"},
			expectedOut: "",
			expectedErr: "Service Principal ID is required (use --id)\n",
		},
		{
			name:        "Claim name is required",
			args:        []string{"application", "claim", "add", "--id", "12345"},
			expectedOut: "",
			expectedErr: "Claim name is required (use --name)\n",
		},
	}

	// Load environment variables
	err := rootcmd.LoadEnv("env_test")
	if err != nil {
		t.Fatalf("Error loading env_test file: %v", err)
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			out, err := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

			// Set command arguments
			rootcmd.RootCmd.SetArgs(tt.args)
			rootcmd.RootCmd.Execute()

			// Assert output and error
			assert.Equal(t, tt.expectedOut, out.String())
			assert.Equal(t, tt.expectedErr, err.String())
		})
	}
}
