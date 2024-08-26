package cmdmanifest

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// manifestGetCmd represents the manifestGet command
var manifestGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application manifest by its AppId",
	Run: func(cmd *cobra.Command, args []string) {
		// originalSelect := clientOptions.Select

		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("Application AppID is required (use --id)")
			return
		}
		// if originalSelect == "" {
		// 	// Default : id,appId,appDisplayName,appOwnerOrganizationId,applicationTemplateId,deletedDateTime,accountEnabled,createdDateTime,displayName,homepage,servicePrincipalNames,signInAudience,tags,appRoleAssignmentRequired,appRoles,oauth2PermissionScopes,preferredSingleSignOnMode,preferredTokenSigningKeyThumbprint,replyUrls,resourceSpecificApplicationPermissions,samlSingleSignOnSettings,servicePrincipalType,tokenEncryptionKeyId
		// 	// Non-default : alternativeNames,addIns,authenticationMethods,certificateBasedAuthConfiguration,certification,claimMappingPolicies,customSecurityAttributes,endpoints,errorUrl,informationalUrls,keyCredentials,licenseDetails,loginUrl,logoutUrl,notes,notificationEmailAddresses,oauth2PermissionGrants,ownedObjects,owners,passThroughUsers,passwordCredentials,publicClient,resourceBehaviorOptions,servicePrincipalLockConfiguration,transitiveMemberOf,verifiedPublisher
		// 	clientOptions.Select = "id,appId,appDisplayName,appOwnerOrganizationId,applicationTemplateId,deletedDateTime,accountEnabled,createdDateTime,displayName,homepage,servicePrincipalNames,signInAudience,tags,appRoleAssignmentRequired,appRoles,oauth2PermissionScopes,preferredSingleSignOnMode,preferredTokenSigningKeyThumbprint,replyUrls,resourceSpecificApplicationPermissions,samlSingleSignOnSettings,servicePrincipalType,tokenEncryptionKeyId,alternativeNames,addIns,authenticationMethods,certificateBasedAuthConfiguration,certification,claimMappingPolicies,customSecurityAttributes,endpoints,errorUrl,informationalUrls,keyCredentials,licenseDetails,loginUrl,logoutUrl,notes,notificationEmailAddresses,oauth2PermissionGrants,ownedObjects,owners,passThroughUsers,passwordCredentials,publicClient,resourceBehaviorOptions,servicePrincipalLockConfiguration,transitiveMemberOf,verifiedPublisher"
		// }
		rootcmd.ClientOptions.Filter = "appId%20eq%20'" + rootcmd.OptID + "'"
		sps, _, err := rootcmd.Client.GetServicePrincipals(rootcmd.ClientOptions)
		apps, _, err := rootcmd.Client.GetApplications(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		if len(apps) != 1 {
			if len(apps) == 0 {
				rootcmd.PrintErrString("No applications found")
			} else {
				rootcmd.PrintErrString("Ambiguous application ID")
			}
			return
		}

		cmd.Println("{\n \"application\": \"" + rootcmd.OutputJSON(apps[0]) + "\",\n \"servicePrincipal\": \"" + rootcmd.OutputJSON(sps[0]) + "\"\n}")
	},
}

func init() {
	manifestCmd.AddCommand(manifestGetCmd)
	// manifestGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
	// 	// Hide flags for this command
	// 	manifestGetCmd.Flags().MarkHidden("batch")
	// 	manifestGetCmd.Flags().MarkHidden("displayname")
	// 	manifestGetCmd.Flags().MarkHidden("post")
	// 	manifestGetCmd.Flags().MarkHidden("search")
	// 	manifestGetCmd.Flags().MarkHidden("select")
	// 	manifestGetCmd.Flags().MarkHidden("skip")
	// 	manifestGetCmd.Flags().MarkHidden("skiptoken")
	// 	manifestGetCmd.Flags().MarkHidden("top")
	// 	// Call parent help func
	// 	command.Parent().HelpFunc()(command, strings)
	// })
}
