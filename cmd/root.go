// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"
	"os"

	"epfl-entra/internal/client"
	httpengine "epfl-entra/internal/client/http_client"
	"epfl-entra/internal/models"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Client is the client to the Entra API
var Client client.Service

// OptBatch is associated with the --batch flag
var OptBatch string

// OptDebug is associated with the --debug flag
var OptDebug bool

// OptDisplayName is associated with the --displayname flag
var OptDisplayName string

// OptEngine is associated with the --engine flag
var OptEngine string

// OptID is associated with the --id flag
var OptID string

// OptPaging is associated with the --paging flag
var OptPaging bool

// OptPostData is associated with the --post flag
var OptPostData string

// OptSearch is associated with the --search flag
var OptSearch string

// OptSelect is associated with the --select flag
var OptSelect string

// OptSkip is associated with the --skip flag
var OptSkip string

// OptSkipToken is associated with the --skiptoken flag
var OptSkipToken string

// OptTop is associated with the --top flag
var OptTop string

// clientOptions is the options (from command line) passed to the client
var clientOptions models.ClientOptions

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ecli",
	Short: "Entra API command line client",
	Long:  `ecli is a command line tool that enables you to interact with Entra API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		if OptEngine == "sdk" {
			panic("SDK engine is not implemented")
			// Client, err = sdkengine.New()
			// if err != nil {
			// 	panic(err)
			// }
		} else {
			if OptDebug {
				fmt.Fprintf(os.Stderr, "ENGINE: %s\n", OptEngine)
			}

			Client, err = httpengine.New()
			if err != nil {
				panic(err)
			}
		}
		if OptDebug {
			fmt.Fprintf(os.Stderr, "Search: %s\n", OptSearch)
			fmt.Fprintf(os.Stderr, "Skip: %s\n", OptSkip)
			fmt.Fprintf(os.Stderr, "Top: %s\n", OptTop)
			fmt.Fprintf(os.Stderr, "Select: %s\n", OptSelect)
		}

		clientOptions = models.ClientOptions{}

		if OptSearch != "" {
			clientOptions.Search = OptSearch
		}

		if OptSelect != "" {
			clientOptions.Select = OptSelect
			clientOptions.Paging = true
		}

		if OptSkip != "" {
			clientOptions.Skip = OptSkip
			clientOptions.Paging = true
		}

		if OptSkipToken != "" {
			clientOptions.SkipToken = OptSkipToken
			clientOptions.Paging = true
		}

		if OptTop != "" {
			clientOptions.Top = OptTop
			clientOptions.Paging = true
		}

		if OptBatch != "" && OptTop == "" {
			clientOptions.Top = OptBatch
			clientOptions.Paging = false
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.epfl-entra.yaml)")
	rootCmd.PersistentFlags().BoolVar(&OptDebug, "debug", false, "Debug mode")
	rootCmd.PersistentFlags().StringVar(&OptBatch, "batch", "900", "Default batch size for client side paging")
	rootCmd.PersistentFlags().StringVar(&OptDisplayName, "displayname", "", "Display name")
	// rootCmd.PersistentFlags().StringVar(&OptEngine, "engine", "rest", "Engine to use ('sdk' or 'rest')")
	rootCmd.PersistentFlags().StringVar(&OptID, "id", "", "Id to use")
	rootCmd.PersistentFlags().StringVar(&OptPostData, "post", "", "Post body data")
	rootCmd.PersistentFlags().StringVar(&OptSearch, "search", "", "Search filter in the form of 'propery:value'")
	rootCmd.PersistentFlags().StringVar(&OptSelect, "select", "", "Comma separated list of properties to be returnded for each object")
	rootCmd.PersistentFlags().StringVar(&OptSkip, "skip", "", "Number of results to skip")
	rootCmd.PersistentFlags().StringVar(&OptSkipToken, "skiptoken", "", "Paging token")
	rootCmd.PersistentFlags().StringVar(&OptTop, "top", "", "Number results to return ('top n results')")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// UGLY!! Find the proper way to do this..
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			_ = godotenv.Load("../../.env")
		}
	}

	tenant := os.Getenv("ENTRA_TENANT")
	if tenant == "" {
		panic("ENTRA_TENANT is not set")
	}

	clientID := os.Getenv("ENTRA_CLIENTID")
	if clientID == "" {
		panic("ENTRA_CLIENTID is not set")
	}

	// Accept empty token (will be retrived by credentials)
}
