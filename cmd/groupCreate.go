// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// groupCreateCmd represents the groupCreate command
var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a group",
	Long: `Create a group whose JSON is passed as argument with --post
	
Example:
  ecli group create --post '{"displayName": "test group AA"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("groupCreate called")
		var group models.Group
		err := json.Unmarshal([]byte(OptPostData), &group)
		if err != nil {
			panic(err)
		}
		err = Client.CreateGroup(&group, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)
}
