package cmdcertificate

import (
	"encoding/json"
	"strings"
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/stretchr/testify/assert"
)

func Test_certificateExpired(t *testing.T) {
	t.Run("Expired returns expired certificates", func(t *testing.T) {
		// Load environment variables
		err := rootcmd.LoadEnv("env_test")
		if err != nil {
			t.Fatalf("Error loading env_test file: %v", err)
		}

		stdout, stderr := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

		rootcmd.RootCmd.SetArgs([]string{"certificate", "expired"})
		rootcmd.RootCmd.Execute()

		outs := strings.Split(stdout.String(), "\n")

		assert.Equal(t, "", stderr.String(), "No error returned")

		// If we have expired certificates, check if first one is valid JSON
		if len(outs) > 0 && outs[0] != "" {
			cert := &models.KeyCredential{}
			certJSON := []byte(outs[0])
			err = json.Unmarshal(certJSON, &cert)
			assert.Nil(t, err, "First expired certificate should be valid JSON")
			assert.True(t, len(cert.KeyID) > 0, "KeyID should be defined")
		}
	})

	t.Run("Expired with end_date filter", func(t *testing.T) {
		// Load environment variables
		err := rootcmd.LoadEnv("env_test")
		if err != nil {
			t.Fatalf("Error loading env_test file: %v", err)
		}

		_, stderr := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

		rootcmd.RootCmd.SetArgs([]string{"certificate", "expired", "--end_date", "2024-12-31"})
		rootcmd.RootCmd.Execute()

		assert.Equal(t, "", stderr.String(), "No error returned")
		// The command should execute successfully even if no certificates match the filter
	})
}
