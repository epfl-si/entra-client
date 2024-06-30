/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
