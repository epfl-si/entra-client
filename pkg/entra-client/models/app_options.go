package models

// AppOptions are use to create an application
type AppOptions struct {
	DisplayName  string
	LogoutURI    string
	MetadataFile string
	RedirectURI  string
	SAMLID       string
	Tags         []string
}
