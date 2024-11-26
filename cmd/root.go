// Package cmd provides the commands for the command line application
package cmd

import (
	"os"

	client "github.com/epfl-si/entra-client/pkg/client"
	httpengine "github.com/epfl-si/entra-client/pkg/client/httpengine"
	"github.com/epfl-si/entra-client/pkg/client/models"

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

// OptPrettyJSON is associated with the --pretty_json flag
var OptPrettyJSON bool

// OptSearch is associated with the --search flag
var OptSearch string

// OptFilter is associated with the --filter flag
var OptFilter string

// OptSelect is associated with the --select flag
var OptSelect string

// OptSkip is associated with the --skip flag
var OptSkip string

// OptSkipToken is associated with the --skiptoken flag
var OptSkipToken string

// OptTop is associated with the --top flag
var OptTop string

// ClientOptions is the options (from command line) passed to the client
var ClientOptions models.ClientOptions

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ecli",
	Short: "Entra API command line client",
	Long:  `ecli is a command line tool that enables you to interact with Entra API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		if OptEngine == "sdk" {
			PrintErrString("SDK engine is not implemented")
			return
			// Client, err = sdkengine.New()
			// if err != nil {
			// 	printErr(err)
			//  return
			// }
		}
		if OptDebug {
			PrintErrString("ENGINE: " + OptEngine + "\n")
		}

		Client, err = httpengine.New()
		if err != nil {
			PrintErr(err)
			return
		}

		ClientOptions = models.ClientOptions{}

		if OptDebug {
			ClientOptions.Debug = true
			PrintErrString("Search: " + OptSearch + "\n")
			PrintErrString("Skip: " + OptSkip + "\n")
			PrintErrString("Top: " + OptTop + "\n")
			PrintErrString("Select: " + OptSelect + "\n")
		}

		if OptSearch != "" {
			ClientOptions.Search = OptSearch
		}

		if OptSelect != "" {
			ClientOptions.Select = OptSelect
			ClientOptions.Paging = true
		}

		if OptSkip != "" {
			ClientOptions.Skip = OptSkip
			ClientOptions.Paging = true
		}

		if OptSkipToken != "" {
			ClientOptions.SkipToken = OptSkipToken
			ClientOptions.Paging = true
		}

		if OptTop != "" {
			ClientOptions.Top = OptTop
			ClientOptions.Paging = true
		}

		if OptBatch != "" && OptTop == "" {
			ClientOptions.Top = OptBatch
			ClientOptions.Paging = false
		}

		if OptFilter != "" && OptTop == "" {
			ClientOptions.Filter = OptFilter
			ClientOptions.Paging = false
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// By default, cmd.Printxxx() functions print to os.Stderr
	RootCmd.SetOut(os.Stdout)

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.epfl-entra.yaml)")
	RootCmd.PersistentFlags().BoolVar(&OptDebug, "debug", false, "Debug mode")
	RootCmd.PersistentFlags().StringVar(&OptBatch, "batch", "900", "Default batch size for client side paging")
	RootCmd.PersistentFlags().StringVar(&OptDisplayName, "displayname", "", "Display name")
	// RootCmd.PersistentFlags().StringVar(&OptEngine, "engine", "rest", "Engine to use ('sdk' or 'rest')")
	RootCmd.PersistentFlags().StringVar(&OptID, "id", "", "Id to use")
	RootCmd.PersistentFlags().StringVar(&OptPostData, "post", "", "Post body data")
	RootCmd.PersistentFlags().BoolVar(&OptPrettyJSON, "pretty_json", false, "JSON pretty output")
	RootCmd.PersistentFlags().StringVar(&OptSearch, "search", "", "Search filter in the form of 'propery:value'")
	RootCmd.PersistentFlags().StringVar(&OptFilter, "filter", "", "Selection filter, ex \" clientId eq 'objectidvalue'\"")
	RootCmd.PersistentFlags().StringVar(&OptSelect, "select", "", "Comma separated list of properties to be returnded for each object")
	RootCmd.PersistentFlags().StringVar(&OptSkip, "skip", "", "Number of results to skip")
	RootCmd.PersistentFlags().StringVar(&OptSkipToken, "skiptoken", "", "Paging token")
	RootCmd.PersistentFlags().StringVar(&OptTop, "top", "", "Number results to return ('top n results')")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

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
		PrintErrString("ENTRA_TENANT is not set")
		return
	}

	clientID := os.Getenv("ENTRA_CLIENTID")
	if clientID == "" {
		PrintErrString("ENTRA_CLIENTID is not set")
		return
	}

	// Accept empty token (will be retrived by credentials)
}
