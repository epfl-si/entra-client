package cmdapplication

import (
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_applicationDelete(t *testing.T) {

	// Transform the function into a test table
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
	}{
		{
			name:        "ID is required",
			args:        []string{"application", "delete"},
			expectedOut: "",
			expectedErr: "ID missing",
		},
	}

	// Load environment variables
	err := rootcmd.LoadEnv("env_test")
	if err != nil {
		t.Fatalf("Error loading env_test file: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset global flags for proper test isolation
			rootcmd.ResetGlobalFlags()

			out, err := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

			rootcmd.RootCmd.SetArgs(tt.args)
			rootcmd.RootCmd.Execute()

			assert.Equal(t, tt.expectedOut, out.String())
			assert.Equal(t, tt.expectedErr, err.String())
		})
	}

}
