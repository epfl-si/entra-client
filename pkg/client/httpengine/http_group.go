package httpengine

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/epfl-si/entra-client/pkg/client/models"

	"io"
)

// CreateGroup creates a group and returns an error
//
// Required permissions: Directory.Read.All
// Required permissions: GroupMember.Read.All
//
// Parameters:
//
//	group: The group to be created
//	opts: The client options
func (c *HTTPClient) CreateGroup(group *models.Group, opts models.ClientOptions) error {
	u, err := json.Marshal(group)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Post("/groups/"+buildQueryString(opts), u, h)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}

// DeleteGroup deletes a group by its Id and returns an error
//
// Required permissions: Group.ReadWrite.All
//
// Parameters:
//
//	id: The group ID
//	opts: The client options
func (c *HTTPClient) DeleteGroup(id string, opts models.ClientOptions) error {
	h := c.buildHeaders(opts)
	response, err := c.RestClient.Delete("/groups/"+id, h)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

// GetGroup returns a group (by its Id) and an error
//
// Required permissions: Directory.Read.All
// Required permissions: GroupMember.Read.All
//
// Parameters:
//
//	id: The group ID
//	opts: The client options
func (c *HTTPClient) GetGroup(id string, opts models.ClientOptions) (*models.Group, error) {
	h := c.buildHeaders(opts)
	response, err := c.RestClient.Get("/groups/"+id+buildQueryString(opts), h)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	var group models.Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// GetGroups gets all groups and returns a slice of groups, a pagination link and an error
//
// Required permissions: Directory.Read.All
// Required permissions: GroupMember.Read.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetGroups(opts models.ClientOptions) ([]*models.Group, string, error) {
	results := make([]*models.Group, 0)

	var response *http.Response
	var err error

	h := c.buildHeaders(opts)

	response, err = c.RestClient.Get("/groups"+buildQueryString(opts), h)

	for {
		if err != nil {
			c.Log.Sugar().Debugf("GetGroups() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		if response.StatusCode != 200 {
			c.Log.Sugar().Debugf("GetGroups() - 1.5 - Status: %s\n", response.Status)
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("GetGroups() - 2 - Error: %s\n", err.Error())
			return nil, "", err
		}

		response.Body.Close()

		var groupResponse models.GroupResponse
		err = json.Unmarshal(body, &groupResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetGroups() - 3 - Error: %s\n", err.Error())
			c.Log.Sugar().Debugf("GetGroups() - 3.5 - Body: %s\n", body)
			return nil, "", err
		}

		results = append(results, groupResponse.Value...)

		if groupResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("GetGroups() - 4 - Calling Next: %s\n", groupResponse.NextLink)
		response, err = c.RestClient.Get(groupResponse.NextLink, h)
		if response.StatusCode != 200 {
			return nil, "", errors.New(response.Status)
		}

		if opts.Paging {
			break
		}
	}

	return results, "", nil
}

// UpdateGroup updates a group and returns an error
//
// Required permissions: ???????
//
// Parameters:
//
//	group: The group to be updated
//	opts: The client options
func (c *HTTPClient) UpdateGroup(group *models.Group, opts models.ClientOptions) error {
	u, err := json.Marshal(group)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	_, err = c.RestClient.Put("/groups/"+group.ID, u, h)
	if err != nil {
		return err
	}

	return nil
}
