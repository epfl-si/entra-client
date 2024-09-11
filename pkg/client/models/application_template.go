// Package models provides the models for the application
package models

// ApplicationTemplate represents an application template
type ApplicationTemplate struct {
	ID                         string   `json:"id,omitempty"`
	DisplayName                string   `json:"displayName,omitempty"`
	HomePageURL                string   `json:"homePageUrl,omitempty"`
	SupportedSingleSignOnModes []string `json:"supportedSingleSignOnModes,omitempty"`
	SupportedProvisioningTypes []string `json:"supportedProvisioningTypes,omitempty"`
	LogoURL                    string   `json:"logoUrl,omitempty"`
	Categories                 []string `json:"categories,omitempty"`
	Publisher                  string   `json:"publisher,omitempty"`
	Description                string   `json:"description,omitempty"`
	ServiceProvisioningErrors  []string `json:"serviceProvisioningErrors,omitempty"`
}
