package models

// IdentityAuthenticationEventListenersIncludeApplicationsBody represents the body
// to add an application to an authenticationEventListeners
type IdentityAuthenticationEventListenersIncludeApplicationsBody struct {
	AppId []string `json:"appId,omitempty"`
}
