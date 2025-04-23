package cmdapplication

import (
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/spf13/cobra"

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
	}

	// Load environment variables
	err := rootcmd.LoadEnv("env_test")
	if err != nil {
		t.Fatalf("Error loading env_test file: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			out, err := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

			rootcmd.RootCmd.SetArgs(tt.args)
			rootcmd.RootCmd.Execute()

			assert.Equal(t, tt.expectedOut, out.String())
			assert.Equal(t, tt.expectedErr, err.String())
		})
	}

}

func Test_init(t *testing.T) {
	// Find the application command in the root command
	var appCmd *cobra.Command
	for _, cmd := range rootcmd.RootCmd.Commands() {
		if cmd.Use == "application" {
			appCmd = cmd
			break
		}
	}
	assert.NotNil(t, appCmd, "Application command should be found in root command")

	// Find the create command in the application command
	var createCmd *cobra.Command
	for _, cmd := range appCmd.Commands() {
		if cmd.Use == "create" {
			createCmd = cmd
			break
		}
	}
	assert.NotNil(t, createCmd, "Create command should be found in application command")

	// Verify that the command has a help function set
	assert.NotNil(t, createCmd.HelpFunc(), "Help function should be set")

	// We need to manually mark the flags as hidden to test the init function's behavior
	// This simulates what happens when the help function is called
	hiddenFlags := []string{"batch", "search", "select", "skip", "skiptoken", "top"}
	for _, flagName := range hiddenFlags {
		// First add the flags if they don't exist
		if createCmd.Flags().Lookup(flagName) == nil {
			createCmd.Flags().String(flagName, "", "")
		}
	}

	// Now call the help function which should hide the flags
	createCmd.HelpFunc()(createCmd, []string{})

	// Now verify that the flags are hidden
	for _, flagName := range hiddenFlags {
		flag := createCmd.Flags().Lookup(flagName)
		if flag != nil {
			assert.True(t, flag.Hidden, "Flag %s should be hidden after help function is called", flagName)
		}
	}
}
