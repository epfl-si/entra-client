package cmdsecret

import (
	"encoding/json"
	"strings"
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/stretchr/testify/assert"
)

func Test_secretList(t *testing.T) {
	t.Run("List returns secrets", func(t *testing.T) {
		// Load environment variables
		err := rootcmd.LoadEnv("env_test")
		if err != nil {
			t.Fatalf("Error loading env_test file: %v", err)
		}

		stdout, stderr := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

		rootcmd.RootCmd.SetArgs([]string{"secret", "list"})
		rootcmd.RootCmd.Execute()

		outs := strings.Split(stdout.String(), "\n")

		assert.True(t, len(outs) >= 1, "At least one secret line returned")
		assert.Equal(t, "", stderr.String(), "No error returned")

		// If we have secrets, check if first one is valid JSON
		if len(outs) > 0 && outs[0] != "" {
			secret := &models.KeyCredential{}
			secretJSON := []byte(outs[0])
			err = json.Unmarshal(secretJSON, &secret)
			assert.Nil(t, err, "First secret should be valid JSON")
			assert.True(t, len(secret.KeyID) > 0, "KeyID should be defined")
		}
	})

	t.Run("List with end_date filter", func(t *testing.T) {
		// Load environment variables
		err := rootcmd.LoadEnv("env_test")
		if err != nil {
			t.Fatalf("Error loading env_test file: %v", err)
		}

		_, stderr := rootcmd.CaptureStdOutputs(rootcmd.RootCmd)

		rootcmd.RootCmd.SetArgs([]string{"secret", "list", "--end_date", "2024-12-31"})
		rootcmd.RootCmd.Execute()

		assert.Equal(t, "", stderr.String(), "No error returned")
		// The command should execute successfully even if no secrets match the filter
	})
}
