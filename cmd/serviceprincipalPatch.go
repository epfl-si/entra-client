// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// serviceprincipalPatchCmd represents the serviceprincipalPatch command
var serviceprincipalPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch a ServicePrincipal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serviceprincipalPatch called")
		var app models.ServicePrincipal
		err := json.Unmarshal([]byte(OptPostData), &app)
		if err != nil {
			panic(err)
		}
		err = Client.PatchServicePrincipal(OptID, &app, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serviceprincipalPatchCmd)
}
