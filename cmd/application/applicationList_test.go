package cmdapplication

import (
	rootcmd "epfl-entra/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applicationList(t *testing.T) {

	// Transform the function into a test table
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
	}{
		{
			name:        "displayname is required",
			args:        []string{"application", "create"},
			expectedOut: "",
			expectedErr: "Name is required (use --displayname)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rout, wout, oldout, rerr, werr, olderr := rootcmd.CaptureOutput()

			rootcmd.RootCmd.SetArgs(tt.args)
			rootcmd.RootCmd.Execute()

			out, err := rootcmd.ReleaseOutput(rout, wout, oldout, rerr, werr, olderr)

			assert.Equal(t, tt.expectedOut, string(out))
			assert.Equal(t, tt.expectedErr, string(err))
		})
	}

}
