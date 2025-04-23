package cmdapplication

import (
	"bytes"
	"testing"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/stretchr/testify/assert"
)

func Test_applicationSecret(t *testing.T) {

	actual := new(bytes.Buffer)
	rootcmd.RootCmd.SetOut(actual)
	// rootCmd.SetErr(actual)
	rootcmd.RootCmd.SetArgs([]string{"application", "secret"})
	rootcmd.RootCmd.Execute()

	expected := "applicationSecret called\n"

	assert.Equal(t, expected, actual.String(), "Generic message displayed")
}
