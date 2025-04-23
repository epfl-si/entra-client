package cmdapplication

import (
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_applicationConsentList(t *testing.T) {

	// Define test cases
	tests := []struct {
		name        string
		args        []string
		expectedOut string
	}{
		{
			name:        "No ID is required",
			args:        []string{"application", "consent", "list"},
			expectedOut: "",
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

			// Capture output
			_, err := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)
			// Set command arguments
			rootcmd.RootCmd.SetArgs(tt.args)
			rootcmd.RootCmd.Execute()

			// Assert output and error
			assert.Equal(t, "", err.String())
		})
	}
}
