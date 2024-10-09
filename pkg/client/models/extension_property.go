// Package models provides the models for the application
package models

// ExtensionProperty represents an extension property
type ExtensionProperty struct {
	AppDisplayName         string   `json:"appDisplayName,omitempty"`
	DataType               string   `json:"dataType,omitempty"`
	DeletedDateTime        string   `json:"deletedDateTime,omitempty"`
	ID                     string   `json:"id,omitempty"`
	IsSyncedFromOnPremises bool     `json:"isSyncedFromOnPremises,omitempty"`
	IsMultiValued          bool     `json:"isMultiValued,omitempty"`
	Name                   string   `json:"name,omitempty"`
	TargetObjects          []string `json:"targetObjects,omitempty"`
}
