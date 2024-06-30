/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// groupGetCmd represents the groupGet command
var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a group by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("groupGet called")
		group, err := Client.GetGroup(OptId, clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Group: %+v\n", group)
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
