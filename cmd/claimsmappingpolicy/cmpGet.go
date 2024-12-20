package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// claimGetCmd represents the claims mapping policy get command
var claimGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a claims mapping policy",
	Long: `This command enables you to get a claims mapping policy.
You can also use the alias "cmp" instead of "claimsmappingpolicy".

	Example:
		./ecli claimsmappingpolicy get --cmpid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
		./ecli cmp get --default
	
`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptCmpID == "" && !OptDefault {
			rootcmd.PrintErrString("Either ID or default is required (use --cmpid or --default)")
			return
		}

		if OptDefault {
			rootcmd.ClientOptions.Default = true
		}

		var cmp *models.ClaimsMappingPolicy
		var err error

		if OptDefault {
			cmps, _, err := rootcmd.Client.GetClaimsMappingPolicies(rootcmd.ClientOptions)
			if err != nil {
				rootcmd.PrintErr(err)
				return
			}

			if len(cmps) != 1 {
				rootcmd.PrintErrString("Default claims mapping policy not found")
				return
			}

			OptCmpID = cmps[0].ID
		}

		cmp, err = rootcmd.Client.GetClaimsMappingPolicy(OptCmpID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Printf("%s\n", rootcmd.OutputJSON(cmp))

	},
}

func init() {
	claimCmd.AddCommand(claimGetCmd)
	// claimGetCmd.MarkFlagsMutuallyExclusive("default", "cmpid")

	claimGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimGetCmd.Flags().MarkHidden("batch")
		claimGetCmd.Flags().MarkHidden("search")
		claimGetCmd.Flags().MarkHidden("select")
		claimGetCmd.Flags().MarkHidden("skip")
		claimGetCmd.Flags().MarkHidden("skiptoken")
		claimGetCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
