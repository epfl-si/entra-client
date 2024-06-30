/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applicationDeleteCmd represents the applicationDelete command
var applicationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationDelete called")
		err := Client.DeleteApplication(OptId, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
