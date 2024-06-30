package sdk_client

import (
	"context"
	"epfl-entra/internal/models"
	"errors"

	"log"

	mm "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func toGroups(groups []mm.Groupable) []*models.Group {

	return []*models.Group{}
}

// CreateGroup creates a group and returns an error
func (c *SDKClient) CreateGroup(group *models.Group, opts models.ClientOptions) error {
	return errors.New("Not implemented")
}

// DeleteGroup deletes a group and returns an error
func (c *SDKClient) DeleteGroup(id string, opts models.ClientOptions) error {
	return errors.New("Not implemented")
}

// GetGroup returns a group and an error
func (c *SDKClient) GetGroup(id string, opts models.ClientOptions) (*models.Group, error) {
	return nil, errors.New("Not implemented")
}

func (c *SDKClient) GetGroups(opts models.ClientOptions) ([]*models.Group, string, error) {
	c.Log.Sugar().Debugf("GetGroups() - 0 - Token: %s\n", c.AccessToken)

	var finalResults []*models.Group
	results, err := c.APIClient.Groups().Get(context.Background(), nil)

	for {
		if err != nil {
			c.Log.Sugar().Debugf("GetGroups() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		addedGroups := toGroups(results.GetValue())
		finalResults = append(finalResults, addedGroups...)

		nextPageURL := results.GetOdataNextLink()
		c.Log.Sugar().Debugf("GetGroups() - 4 - Calling Next: %s\n", *nextPageURL)
		if nextPageURL != nil {
			results, err = c.APIClient.Groups().WithUrl(*nextPageURL).Get(context.Background(), nil)
			if err != nil {
				log.Fatalf("Error getting messages: %v\n", err)
			}
		} else {
			break
		}
	}

	return finalResults, "", nil
}
