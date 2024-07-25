// Package cmd provides the commands for the command line application
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
		group, err := Client.GetGroup(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Group: %s\n", OutputJSON(group))
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)
}
