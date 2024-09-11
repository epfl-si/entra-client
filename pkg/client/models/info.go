package models

// Info represents the info of a service principal
type Info struct {
	LogoURL             string `json:"logoUrl,omitempty"`
	MarketingURL        string `json:"marketingUrl,omitempty"`
	PrivacyStatementURL string `json:"privacyStatementUrl,omitempty"`
	SupportURL          string `json:"supportUrl,omitempty"`
	TermsOfServiceURL   string `json:"termsOfServiceUrl,omitempty"`
}
