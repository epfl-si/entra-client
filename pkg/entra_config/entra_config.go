// package entraconfig provides the configuration for the Entra API
package entraconfig

var config = map[string]map[string]string{
	"b6cddbc1-2348-4644-af0a-2fdb55573e3b": { // Test tenant
		"MICROSOFT_GRAPH_API_ID": "ab3b94f2-841b-4ca3-8f3e-7e63b5a5e233",
	},
	"f6c2556a-c4fb-4ab1-a2c7-9e220df11c43": { // Prod tenant
		"MICROSOFT_GRAPH_API_ID": "",
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
func (c *EntraConfig) Get(key string) string {
	// test if key exists
	val, ok := c.Value[key]
	if !ok {
		panic("Key " + key + " does not exist")
	}
	return val
}
