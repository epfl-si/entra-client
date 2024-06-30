package models

type User struct {
	ID                string   `json:"id,omitempty"`
	BusinessPhones    []string `json:"businessPhones,omitempty"`
	DisplayName       string   `json:"displayName,omitempty"`
	GivenName         string   `json:"givenName,omitempty"`
	JobTitle          string   `json:"jobTitle,omitempty"`
	Mail              string   `json:"mail,omitempty"`
	MobilePhone       string   `json:"mobilePhone,omitempty"`
	OfficeLocation    string   `json:"officeLocation,omitempty"`
	PreferredLanguage string   `json:"preferredLanguage,omitempty"`
	Surname           string   `json:"surname,omitempty"`
	UserPrincipalName string   `json:"userPrincipalName,omitempty"`
}
