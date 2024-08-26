package cmdapplication

import (
	"bytes"
	rootcmd "epfl-entra/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AnyCommandThroughRootCmd(t *testing.T) {

	actual := new(bytes.Buffer)
	rootcmd.RootCmd.SetOut(actual)
	// rootCmd.SetErr(actual)
	rootcmd.RootCmd.SetArgs([]string{"application"})
	rootcmd.RootCmd.Execute()

	expected := "application called\n"

	assert.Equal(t, expected, actual.String(), "placeholder message expected")
}
