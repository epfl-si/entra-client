package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userGetCmd represents the userGet command
var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("userGet called")
		user, err := Client.GetUser(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Printf("User: %+v\n", user)
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
