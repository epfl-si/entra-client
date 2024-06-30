/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"epfl-entra/internal/client"
	"epfl-entra/internal/client/http_client"
	"epfl-entra/internal/client/sdk_client"
	"epfl-entra/internal/models"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var EntraClient *http_client.HTTPClient

var Client client.Service
var OptBatch string
var OptDebug bool
var OptEngine string
var OptId string
var OptPaging bool
var OptPostData string
var OptSearch string
var OptSelect string
var OptSkip string
var OptSkipToken string
var OptTop string

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
			Client, err = sdk_client.New()
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("ENGINE: %s\n", OptEngine)

			Client, err = http_client.New()
			if err != nil {
				panic(err)
			}
		}
		fmt.Printf("Top: %s\n", OptTop)
		fmt.Printf("Select: %s\n", OptSelect)

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
	var err error

	err = rootCmd.Execute()
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
	rootCmd.PersistentFlags().StringVar(&OptEngine, "engine", "rest", "Engine to use ('sdk' or 'rest')")
	rootCmd.PersistentFlags().StringVar(&OptId, "id", "", "Id to use")
	rootCmd.PersistentFlags().StringVar(&OptPostData, "post", "", "Post body data")
	rootCmd.PersistentFlags().StringVar(&OptSearch, "search", "", "Search filter in the form of 'propery:value'")
	rootCmd.PersistentFlags().StringVar(&OptSelect, "select", "", "Comma separated list of properties to be returnded for each object")
	rootCmd.PersistentFlags().StringVar(&OptSkip, "skip", "", "Number of results to skip")
	rootCmd.PersistentFlags().StringVar(&OptSkipToken, "skiptoken", "", "Paging token")
	rootCmd.PersistentFlags().StringVar(&OptTop, "top", "", "Number results to return ('top n results')")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	godotenv.Load()

	secret := os.Getenv("ENTRA_SECRET")
	if secret == "" {
		panic("ENTRA_SECRET is not set")
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

func debug(msg string) {
	if OptDebug {
		fmt.Println(msg)
	}
}

func debugf(format string, args ...interface{}) {
	if OptDebug {
		fmt.Printf(format, args...)
	}
}

func OutputJSON(data interface{}) string {
	jdata, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(jdata)
}
