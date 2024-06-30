/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applicationList.goCmd represents the applicationList.go command
var applicationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		applications, _, err := Client.GetApplications(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, application := range applications {
			fmt.Printf("%s\n", OutputJSON(application))
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationList.goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationList.goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
