package httpengine

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/epfl-si/entra-client/pkg/client/models"
)

// GetOAuth2PermissionGrants returns a list of OAuth2 permission grants
//
// Resources: https://learn.microsoft.com/en-us/graph/api/oauth2permissiongrant-list?view=graph-rest-1.0
func (c *HTTPClient) GetOAuth2PermissionGrants(opts models.ClientOptions) ([]*models.OAuth2PermissionGrant, string, error) {
	results := make([]*models.OAuth2PermissionGrant, 0)
	var grantResponse models.OAuth2PermissionGrantResponse

	h := c.buildHeaders(opts)
	url := "/oauth2PermissionGrants" + buildQueryString(opts)

	if opts.Debug {
		c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - Calling: %s\n", url)
	}

	response, err := c.RestClient.Get(url, h)
	if err != nil {
		c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - REST Error: %s\n", err.Error())
		return nil, "", err
	}
	defer response.Body.Close()

	if opts.Debug {
		c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - Response Status: %d\n", response.StatusCode)
	}

	page := 1
	for {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - ReadAll Error: %s\n", err.Error())
			return nil, "", err
		}

		if opts.Debug {
			c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - Page %d - Response Body: %s\n", page, string(body))
		}

		// Reset grantResponse to avoid retaining NextLink from previous iteration
		grantResponse = models.OAuth2PermissionGrantResponse{}

		err = json.Unmarshal(body, &grantResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - Unmarshal Error: %s - Body: %s\n", err.Error(), string(body))
			return nil, "", err
		}

		response.Body.Close()

		if opts.Debug {
			c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - Page %d - Found %d grants\n", page, len(grantResponse.Value))
		}

		for i := range grantResponse.Value {
			results = append(results, &grantResponse.Value[i])
		}

		if grantResponse.NextLink == "" {
			if opts.Debug {
				c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - No more pages, total grants: %d\n", len(results))
			}
			break
		}

		if opts.Debug {
			c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - Page %d - NextLink: %s\n", page, grantResponse.NextLink)
		}

		response, err = c.RestClient.Get(grantResponse.NextLink, h)
		if err != nil {
			c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - NextLink Error: %s\n", err.Error())
			return nil, "", err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - NextLink Status Error: %d - %s\n", response.StatusCode, response.Status)
			return nil, "", errors.New(response.Status)
		}

		if opts.Paging {
			if opts.Debug {
				c.Log.Sugar().Debugf("GetOAuth2PermissionGrants() - Paging mode, stopping after page %d\n", page)
			}
			break
		}

		page++
	}

	return results, grantResponse.NextLink, nil
}
