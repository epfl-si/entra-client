package models

// ParentalControlSettings represents the parental control settings
type ParentalControlSettings struct {
	CountriesBlockedForMinors []string `json:"countriesBlockedForMinors,omitempty"`
	LegalAgeGroupRule         string   `json:"legalAgeGroupRule,omitempty"`
}
