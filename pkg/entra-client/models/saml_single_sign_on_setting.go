package models

// SAMLSingleSignOnSetting represents a SAML single sign-on setting
type SAMLSingleSignOnSetting struct {
	RelayState string `json:"relayState,omitempty"`
}
