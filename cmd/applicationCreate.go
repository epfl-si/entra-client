// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// applicationCreateCmd represents the applicationCreate command
var applicationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an application",
	Long: `Create an application whose JSON is passed as argument with --post
	
Example:
  ecli application create --engine sdk --post '{"displayName": "test API POST AA"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		var app models.Application
		err := json.Unmarshal([]byte(OptPostData), &app)
		if err != nil {
			panic(err)
		}
		fmt.Println("applicationCreate called")
		err = Client.CreateApplication(&app, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationCreateCmd)
}
