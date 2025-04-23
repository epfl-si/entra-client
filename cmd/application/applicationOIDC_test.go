package cmdapplication

import (
	"bytes"
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_applicationOIDC(t *testing.T) {

	actual := new(bytes.Buffer)
	rootcmd.RootCmd.SetOut(actual)
	// rootCmd.SetErr(actual)
	rootcmd.RootCmd.SetArgs([]string{"application", "oidc"})
	rootcmd.RootCmd.Execute()

	expected := "applicationOIDC called\n"

	assert.Equal(t, expected, actual.String(), "placeholder message expected")
}
