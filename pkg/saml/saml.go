// Package saml provides the SAML functionality for the application
package saml

import (
	"os"

	samlorigin "github.com/crewjam/saml"
	"github.com/crewjam/saml/samlsp"
)

// EntityDescriptor represents the SAML EntityDescriptor object.
//
// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.3.2
// type EntityDescriptor struct {
// 	XMLName                       xml.Name      `xml:"urn:oasis:names:tc:SAML:2.0:metadata EntityDescriptor"`
// 	EntityID                      string        `xml:"entityID,attr"`
// 	ID                            string        `xml:",attr,omitempty"`
// 	ValidUntil                    time.Time     `xml:"validUntil,attr,omitempty"`
// 	CacheDuration                 time.Duration `xml:"cacheDuration,attr,omitempty"`
// 	Signature                     *etree.Element
// 	RoleDescriptors               []RoleDescriptor               `xml:"RoleDescriptor"`
// 	IDPSSODescriptors             []IDPSSODescriptor             `xml:"IDPSSODescriptor"`
// 	SPSSODescriptors              []SPSSODescriptor              `xml:"SPSSODescriptor"`
// 	AuthnAuthorityDescriptors     []AuthnAuthorityDescriptor     `xml:"AuthnAuthorityDescriptor"`
// 	AttributeAuthorityDescriptors []AttributeAuthorityDescriptor `xml:"AttributeAuthorityDescriptor"`
// 	PDPDescriptors                []PDPDescriptor                `xml:"PDPDescriptor"`
// 	AffiliationDescriptor         *AffiliationDescriptor
// 	Organization                  *Organization
// 	ContactPerson                 *ContactPerson
// 	AdditionalMetadataLocations   []string `xml:"AdditionalMetadataLocation"`
// }
// SSODescriptor represents the SAML complex type SSODescriptor
//
// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.4.2
// type SSODescriptor struct {
// 	RoleDescriptor
// 	ArtifactResolutionServices []IndexedEndpoint `xml:"ArtifactResolutionService"`
// 	SingleLogoutServices       []Endpoint        `xml:"SingleLogoutService"`
// 	ManageNameIDServices       []Endpoint        `xml:"ManageNameIDService"`
// 	NameIDFormats              []NameIDFormat    `xml:"NameIDFormat"`
// }

// IDPSSODescriptor represents the SAML IDPSSODescriptorType object.
//
// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.4.3
// type IDPSSODescriptor struct {
// 	XMLName xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:metadata IDPSSODescriptor"`
// 	SSODescriptor
// 	WantAuthnRequestsSigned *bool `xml:",attr"`

// 	SingleSignOnServices       []Endpoint  `xml:"SingleSignOnService"`
// 	ArtifactResolutionServices []Endpoint  `xml:"ArtifactResolutionService"`
// 	NameIDMappingServices      []Endpoint  `xml:"NameIDMappingService"`
// 	AssertionIDRequestServices []Endpoint  `xml:"AssertionIDRequestService"`
// 	AttributeProfiles          []string    `xml:"AttributeProfile"`
// 	Attributes                 []Attribute `xml:"Attribute"`
// }

// SPSSODescriptor represents the SAML SPSSODescriptorType object.
//
// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.4.2
// // type SPSSODescriptor struct {
// // 	XMLName xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:metadata SPSSODescriptor"`
// // 	SSODescriptor
// // 	AuthnRequestsSigned        *bool                       `xml:",attr"`
// // 	WantAssertionsSigned       *bool                       `xml:",attr"`
// // 	AssertionConsumerServices  []IndexedEndpoint           `xml:"AssertionConsumerService"`
// // 	AttributeConsumingServices []AttributeConsumingService `xml:"AttributeConsumingService"`
// // }

// IndexedEndpoint represents the SAML IndexedEndpointType object.
//
// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.2.3
// type IndexedEndpoint struct {
// 	Binding          string  `xml:"Binding,attr"`
// 	Location         string  `xml:"Location,attr"`
// 	ResponseLocation *string `xml:"ResponseLocation,attr,omitempty"`
// 	Index            int     `xml:"index,attr"`
// 	IsDefault        *bool   `xml:"isDefault,attr"`
// }

// Endpoint represents the SAML EndpointType object.
//
// See http://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf §2.2.2
//
//	type Endpoint struct {
//		Binding          string `xml:"Binding,attr"`
//		Location         string `xml:"Location,attr"`
//		ResponseLocation string `xml:"ResponseLocation,attr,omitempty"`
//	}

type EntityDescriptor struct {
	samlorigin.EntityDescriptor
}

// GetMetadata returns the metadata for the SAML XML metadata file whose name is passed as argument
func GetMetadata(fileName string) (*EntityDescriptor, error) {
	xmlFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	metadata, err := samlsp.ParseMetadata(xmlFile)
	if err != nil {
		return nil, err
	}

	return &EntityDescriptor{*metadata}, nil
}
