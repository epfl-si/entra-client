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

	expected_out := "application called\n"
	expected_err := ""

	assert.Equal(t, expected_out, out.String(), "actual is not expected")
	assert.Equal(t, expected_err, err.String(), "actual is not expected")
}
