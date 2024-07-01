// Package cmd provides the command line application for the application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// userCreateCmd represents the userCreate command
var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Long: `Create a user whose JSON is passed as argument with --post
	
Example:
  ecli user create --post '{"displayName": "test user AA"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("userCreate called")
		var app models.User
		err := json.Unmarshal([]byte(OptPostData), &app)
		if err != nil {
			panic(err)
		}
		err = Client.CreateUser(&app, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
