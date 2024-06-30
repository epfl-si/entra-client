/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applicationGetCmd represents the applicationGet command
var applicationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationGet called")
		application, err := Client.GetApplication(OptId, clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Application: %+v\n", application)
	},
}

func init() {
	applicationCmd.AddCommand(applicationGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
