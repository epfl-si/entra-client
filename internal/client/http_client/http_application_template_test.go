package httpengine

import (
	"epfl-entra/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApplicationTemplate(t *testing.T) {
	// This is a placeholder and should be replaced with actual setup, execution and assertion
	t.Run("test case 1", func(t *testing.T) {
		client := &HTTPClient{}
		opts := models.ClientOptions{}
		result, err := client.GetApplicationTemplate("testID", opts)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
	// Add more test cases as needed
}

func TestGetApplicationTemplates(t *testing.T) {
	// This is a placeholder and should be replaced with actual setup, execution and assertion
	t.Run("test case 1", func(t *testing.T) {
		client := &HTTPClient{}
		opts := models.ClientOptions{}
		result, nextLink, err := client.GetApplicationTemplates(opts)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "", nextLink)
	})
	// Add more test cases as needed
}
