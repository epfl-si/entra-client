package cmdapplication

import (
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_applicationUserList(t *testing.T) {

	// Define test cases
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
	}{
		{
			name:        "Service Principal or Application ID is required",
			args:        []string{"application", "user", "list"},
			expectedOut: "",
			expectedErr: "Service Principal or application ID is required (use --spid or --appid)\n",
		},
		{
			name:        "Service Principal or Application ID is required",
			args:        []string{"application", "user", "list", "--appid", "12345"},
			expectedOut: "",
			expectedErr: "No service principal found for this appID: 400 Bad Request\n",
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
			stdOut, stdErr := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

			// Set command arguments
			rootcmd.RootCmd.SetArgs(tt.args)
			rootcmd.RootCmd.Execute()

			// Assert output and error
			assert.Equal(t, tt.expectedOut, stdOut.String())
			assert.Equal(t, tt.expectedErr, stdErr.String())
		})
	}
}
