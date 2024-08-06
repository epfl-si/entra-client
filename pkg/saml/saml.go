package saml

import (
	"os"

	"github.com/crewjam/saml"
	"github.com/crewjam/saml/samlsp"
)

// GetMetadata returns the metadata for the SAML XML metadata file whose name is passed as argument
func GetMetadata(fileName string) (*saml.EntityDescriptor, error) {
	xmlFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	metadata, err := samlsp.ParseMetadata(xmlFile)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}
