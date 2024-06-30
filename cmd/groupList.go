/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// groupListCmd represents the groupList command
var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Run: func(cmd *cobra.Command, args []string) {
		groups, _, err := Client.GetGroups(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, group := range groups {
			fmt.Printf("%s\n", OutputJSON(group))
		}
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
