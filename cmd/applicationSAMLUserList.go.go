// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// applicationSAMLUserList.goCmd represents the applicationSAMLUserList.go command
var applicationSAMLUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users for a SAML application given its service principal ID",
	Run: func(cmd *cobra.Command, args []string) {
		sp, err := Client.GetServicePrincipal(OptID, clientOptions)
		if err != nil {
			panic(err)
		}
		for _, user := range sp.AppRoles {
			cmd.Println(OutputJSON(user))
		}
	},
}

func init() {
	applicationSAMLUserCmd.AddCommand(applicationSAMLUserListCmd)
}
