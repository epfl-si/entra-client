package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AnyCommandThroughRootCmd(t *testing.T) {

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	// rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"application"})
	rootCmd.Execute()

	expected := "application called\n"

	assert.Equal(t, expected, actual.String(), "placeholder message expected")
}
