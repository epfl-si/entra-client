package models

// OptionalClaims holds optional claims definition
type OptionalClaims struct {
	AccessToken []OptionalClaim `json:"accessToken,omitempty"`
	IDToken     []OptionalClaim `json:"idToken,omitempty"`
	SAML2Token  []OptionalClaim `json:"saml2Token,omitempty"`
}
