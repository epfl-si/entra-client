package cmdcertificate

import (
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_certificate(t *testing.T) {
	t.Run("Certificate command displays help", func(t *testing.T) {
		// Load environment variables
		err := rootcmd.LoadEnv("env_test")
		if err != nil {
			t.Fatalf("Error loading env_test file: %v", err)
		}

		stdout, stderr := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

		rootcmd.RootCmd.SetArgs([]string{"certificate"})
		rootcmd.RootCmd.Execute()

		assert.Contains(t, stdout.String(), "certificate called")
		assert.Equal(t, "", stderr.String(), "No error returned")
	})
}
