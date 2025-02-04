package cmdapplication

import (
	"encoding/json"
	"strings"
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

func Test_applicationList(t *testing.T) {
	t.Run("List returns multiple application", func(t *testing.T) {
		godotenv.Load("../../.env")
		rout, wout, oldout, rerr, werr, olderr := rootcmd.CaptureOutput()

		rootcmd.RootCmd.SetArgs([]string{"application", "list"})
		rootcmd.RootCmd.Execute()

		stdout, stderr := rootcmd.ReleaseOutput(rout, wout, oldout, rerr, werr, olderr)

		outs := strings.Split(string(stdout), "\n")

		assert.True(t, len(outs) > 1, "More than one application returned")
		assert.Equal(t, "", string(stderr))

		// unmarshal out[0] to check if it is a valid JSON
		app := &models.Application{}
		appJSON := []byte(outs[0])
		err := json.Unmarshal(appJSON, &app)
		assert.Nil(t, err)
		assert.True(t, len(app.ID) > 1, "ID defined ("+app.ID+")")
		assert.True(t, len(app.AppID) > 1, "AppID defined")
		assert.True(t, len(app.DisplayName) > 1, "DisplayName defined")
	})

}
