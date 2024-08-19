package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applicationCreate(t *testing.T) {

	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	rootCmd.SetOut(out)
	rootCmd.SetErr(err)
	rootCmd.SetArgs([]string{"application", "create"})
	rootCmd.Execute()

	expectedOut := "application called\n"
	expectedErr := ""

	assert.Equal(t, expectedOut, out.String(), "actual is not expected")
	assert.Equal(t, expectedErr, err.String(), "actual is not expected")
}
