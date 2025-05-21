// package entraconfig provides the configuration for the Entra API
package entraconfig

var config = map[string]map[string]string{
	"b6cddbc1-2348-4644-af0a-2fdb55573e3b": { // Test tenant
		"MSGRAPH_API_ID":                     "ab3b94f2-841b-4ca3-8f3e-7e63b5a5e233",
		"MSGRAPH_API_RESOURCE_APP_ID":        "00000003-0000-0000-c000-000000000000",
		"MSGRAPH_EMAIL_RESOURCE_ID":          "64a6cdd6-aab1-4aaf-94b8-3cc8405e90d0",
		"MSGRAPH_OFFLINE_ACCESS_RESOURCE_ID": "7427e0e9-2fba-42fe-b0c0-848c9e6a8182",
		"MSGRAPH_OPENID_RESOURCE_ID":         "37f7f235-527c-4136-accd-4a02d197296e",
		"MSGRAPH_PROFILE_RESOURCE_ID":        "14dad69e-099b-42c9-810b-d002981feec1",
		"MSGRAPH_USER_READ_RESOURCE_ID":      "e1fe6dd8-ba31-4d61-89e7-88639da4683d",
		"AAD_All Staff Users_ID":             "ecd361a9-0089-451d-b851-a4223aad73f7",
		"AAD_All Student Users_ID":           "1f8006b6-ae21-42de-957c-7487cdbe7ddd",
		"AAD_All Hosts Users_ID":             "02ab45fd-6c06-4053-8aa9-06068929d806",
		"AAD_All Outside EPFL Users":         "43c1e1df-2a86-44c4-abba-9656aeeac56d",

		// Claim Mapping Policies
		"CMP_BASE_ID":            "958388be-03aa-419e-9dba-36d05b6a3fd5",
		"CMP_CFS_ID":             "9ba12324-5f7a-4ce6-9886-1df43eaa948a",
		"CMP_AUTH_ID":            "57511031-0bfc-49ca-8c3f-85db7cf86301",
		"CMP_CFS_AUTH_ID":        "1f15daa6-1b68-4719-ab1e-5db3f01641ed",
		"CMP_ACCRED_ID":          "4157bad6-de2d-44e7-9367-64362eefd44a",
		"CMP_CFS_ACCRED_ID":      "4ac6d1d4-ef60-462f-a8a8-eff176ba130a",
		"CMP_AUTH_ACCRED_ID":     "4dd8108c-030d-411e-b084-31de832521b6",
		"CMP_CFS_AUTH_ACCRED_ID": "518db959-c595-419a-878f-a0681a329a8b",
	},
	"f6c2556a-c4fb-4ab1-a2c7-9e220df11c43": { // Prod tenant
		"MSGRAPH_API_ID":                     "0b24c62a-4a9d-4b87-aeeb-ad5343e4d7ca",
		"MSGRAPH_API_RESOURCE_APP_ID":        "00000003-0000-0000-c000-000000000000",
		"MSGRAPH_EMAIL_RESOURCE_ID":          "64a6cdd6-aab1-4aaf-94b8-3cc8405e90d0",
		"MSGRAPH_OFFLINE_ACCESS_RESOURCE_ID": "7427e0e9-2fba-42fe-b0c0-848c9e6a8182",
		"MSGRAPH_OPENID_RESOURCE_ID":         "37f7f235-527c-4136-accd-4a02d197296e",
		"MSGRAPH_PROFILE_RESOURCE_ID":        "14dad69e-099b-42c9-810b-d002981feec1",
		"MSGRAPH_USER_READ_RESOURCE_ID":      "e1fe6dd8-ba31-4d61-89e7-88639da4683d",
		"AAD_All Staff Users_ID":             "6966acae-2e9a-456f-839c-79e2e00a5ade",
		"AAD_All Student Users_ID":           "ae4dcfd5-d53d-4f80-8847-3dbc6a488a37",
		"AAD_All Hosts Users_ID":             "f18bc55b-621a-4ac1-8250-1226b4002793",
		"AAD_All Outside EPFL Users_ID":      "8689a95a-897b-4b88-8ae6-7b83979a93f0",

		// Claim Mapping Policies
		"CMP_BASE_ID":            "d52eefc3-d469-4a14-b210-2bd104d54164",
		"CMP_CFS_ID":             "5ab4b8b5-530c-4fe0-b356-235a601e8940",
		"CMP_AUTH_ID":            "f63a0e2b-7942-4c54-9492-b46b5fa6564e",
		"CMP_CFS_AUTH_ID":        "65d071d9-d745-4b43-88a9-bac58f9cf373",
		"CMP_ACCRED_ID":          "a9a519a7-1026-44cf-ae3d-c0dd76db4e83",
		"CMP_CFS_ACCRED_ID":      "0ba5a10b-815b-4056-92db-6f24ac79a856",
		"CMP_AUTH_ACCRED_ID":     "49a22fb2-fe20-4f3b-a2c8-ce438da717b5",
		"CMP_CFS_AUTH_ACCRED_ID": "13f1b5ae-e77f-4606-a775-71c871dc5cae",
	},
}

// EntraConfig is a struct that holds the configuration for a given tenant
type EntraConfig struct {
	TenantID string
	Value    map[string]string
}

// New creates a new EntraConfig object for the given tenant ID
func New(tenantID string) *EntraConfig {
	// TODO: could check here that all keys in every tenants are defined in all tenants
	config, ok := config[tenantID]
	if !ok {
		panic("Tenant ID " + tenantID + " does not exist")
	}
	return &EntraConfig{TenantID: tenantID, Value: config}
}

// Get return the value of a key for the Entra configuration in the current tenant
// TODO: look in COMMON_CONFIG if the key is not found in the current tenant
// (enable to have a common configuration for all tenants: MS Graph API ID, etc.)
func (c *EntraConfig) Get(key string) string {
	// test if key exists
	val, ok := c.Value[key]
	if !ok {
		panic("Key " + key + " does not exist")
	}
	return val
}

func (c *EntraConfig) Has(key string) bool {
	_, ok := c.Value[key]
	return ok
}
