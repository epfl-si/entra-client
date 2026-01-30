package utils

import (
	"fmt"
	"testing"
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
			got, err := NormalizeName(tt.name, tt.env)
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

func TestNormalizeURI(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"remove trailing slash", "https://example.com/", "https://example.com"},
		{"no trailing slash", "https://example.com", "https://example.com"},
		{"convert http to https", "http://example.com", "https://example.com"},
		{"convert http to https with trailing slash", "http://example.com/", "https://example.com"},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeURI(tt.input)
			if got != tt.want {
				t.Errorf("NormalizeURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeThumbprint(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"with spaces", "12 34 56 78", "12345678"},
		{"with dashes", "12-34-56-78", "12345678"},
		{"with spaces and dashes", "12 34-56 78", "12345678"},
		{"no spaces or dashes", "12345678", "12345678"},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeThumbprint(tt.input)
			if got != tt.want {
				t.Errorf("NormalizeThumbprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
