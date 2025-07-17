package cmdsecret

import (
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_secret(t *testing.T) {
	t.Run("Secret command displays help", func(t *testing.T) {
		// Load environment variables
		err := rootcmd.LoadEnv("env_test")
		if err != nil {
			t.Fatalf("Error loading env_test file: %v", err)
		}

		stdout, stderr := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

		rootcmd.RootCmd.SetArgs([]string{"secret"})
		rootcmd.RootCmd.Execute()

		assert.Contains(t, stdout.String(), "secret called")
		assert.Equal(t, "", stderr.String(), "No error returned")
	})
}
