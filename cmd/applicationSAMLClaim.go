/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var OptClaimPolicyID string

// applicationSAMLClaimCmd represents the applicationSAMLClaim command
var applicationSAMLClaimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Manage claims mapping for a SAML application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLClaim called")
	},
}

func init() {
	applicationSAMLCmd.AddCommand(applicationSAMLClaimCmd)

	applicationSAMLClaimCmd.PersistentFlags().StringVar(&OptClaimPolicyID, "claimpolicyid", "", "The claim policy ID")
}
