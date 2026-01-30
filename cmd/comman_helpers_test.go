package cmd

import (
	"fmt"
	"testing"

	"github.com/epfl-si/entra-client/pkg/utils"
)

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		name        string
		env         int
		want        string
		expectedErr error
	}{
		{"EPFL - test", 1, "EPFL - test - DEV", nil},
		{"EPFL - test - DEV", 2, "EPFL - test - TEST", nil},
		{"EPFL - test - TEST", 3, "EPFL - test", nil},
		{"test", 1, "EPFL - test - DEV", nil},
		{"test", 2, "EPFL - test - TEST", nil},
		{"test", 3, "EPFL - test", nil},
		{"test", 4, "EPFL - test", fmt.Errorf("invalid environment: %d, must be 1, 2 or 3", 4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.NormalizeName(tt.name, tt.env)
			if tt.expectedErr != nil {
				if err == nil {
					// If we expected an error but got nil, fail the test
					t.Errorf("NormalizeName() error = nil, want %v", tt.expectedErr)
					return
				}
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("NormalizeName() error = %v, want %v", err, tt.expectedErr)
					return
				}
				return // We got the expected error
			}
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("NormalizeName() error = %v, want %v", err, tt.expectedErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
