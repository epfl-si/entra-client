package models

import "time"

type Application struct {
	ID                            string     `json:"id,omitempty"`
	DeletedDateTime               *time.Time `json:"deletedDateTime,omitempty"`
	Classification                *string    `json:"classification,omitempty"`
	CreatedDateTime               time.Time  `json:"createdDateTime,omitempty"`
	CreationOptions               []string   `json:"creationOptions,omitempty"`
	Description                   *string    `json:"description,omitempty"`
	DisplayName                   string     `json:"displayName,omitempty"`
	ExpirationDateTime            *time.Time `json:"expirationDateTime,omitempty"`
	GroupTypes                    []string   `json:"groupTypes,omitempty"`
	IsAssignableToRole            *bool      `json:"isAssignableToRole,omitempty"`
	Mail                          *string    `json:"mail,omitempty"`
	MailEnabled                   bool       `json:"mailEnabled,omitempty"`
	MailNickname                  string     `json:"mailNickname,omitempty"`
	MembershipRule                *string    `json:"membershipRule,omitempty"`
	MembershipRuleProcessingState *string    `json:"membershipRuleProcessingState,omitempty"`
	OnPremisesDomainName          string     `json:"onPremisesDomainName,omitempty"`
	OnPremisesLastSyncDateTime    time.Time  `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesNetBiosName         string     `json:"onPremisesNetBiosName,omitempty"`
	OnPremisesSamAccountName      string     `json:"onPremisesSamAccountName,omitempty"`
	OnPremisesSecurityIdentifier  string     `json:"onPremisesSecurityIdentifier,omitempty"`
	OnPremisesSyncEnabled         bool       `json:"onPremisesSyncEnabled,omitempty"`
	PreferredDataLocation         *string    `json:"preferredDataLocation,omitempty"`
	PreferredLanguage             *string    `json:"preferredLanguage,omitempty"`
	ProxyAddresses                []string   `json:"proxyAddresses,omitempty"`
	RenewedDateTime               time.Time  `json:"renewedDateTime,omitempty"`
	ResourceBehaviorOptions       []string   `json:"resourceBehaviorOptions,omitempty"`
	ResourceProvisioningOptions   []string   `json:"resourceProvisioningOptions,omitempty"`
	SecurityEnabled               bool       `json:"securityEnabled,omitempty"`
	SecurityIdentifier            string     `json:"securityIdentifier,omitempty"`
	Theme                         *string    `json:"theme,omitempty"`
	UniqueName                    *string    `json:"uniqueName,omitempty"`
	Visibility                    *string    `json:"visibility,omitempty"`
	OnPremisesProvisioningErrors  []string   `json:"onPremisesProvisioningErrors,omitempty"`
	ServiceProvisioningErrors     []string   `json:"serviceProvisioningErrors,omitempty"`
}
