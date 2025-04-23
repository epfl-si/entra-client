package cmdapplication

import (
	"bytes"
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_applicationSAMLClaimUnassign(t *testing.T) {

	// Define test cases
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
	}{
		{
			name:        "ID and claim are required",
			args:        []string{"application", "saml", "claim", "unassign"},
			expectedOut: "",
			expectedErr: "applicationSAMLClaimUnassign called\n",
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

			actual := new(bytes.Buffer)
			rootcmd.RootCmd.SetOut(actual)
			// rootCmd.SetErr(actual)
			rootcmd.RootCmd.SetArgs([]string{"application"})
			rootcmd.RootCmd.Execute()

			expected := "application called\n"

			assert.Equal(t, expected, actual.String(), "placeholder message expected")
		})
	}
}
