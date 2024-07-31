// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serviceprincipalGetCmd represents the serviceprincipalGet command
var serviceprincipalGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a service principal by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if OptID == "" {
			panic("Service Principal ID is required (use --id)")
		}
		user, err := Client.GetServicePrincipal(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", OutputJSON(user))
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceprincipalGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceprincipalGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
