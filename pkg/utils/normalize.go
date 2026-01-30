package utils

import (
	"fmt"
	"regexp"
)

// NormalizeName normalizes a name by adding prefixing with EPFL and suffixing with environment when it's not PROD, if it's not already the case
// Note: " - " is replaced by "-" in the name.
//   - name: the name to normalize
//   - env: the environment (1: DEV, 2: TEST, 3: PROD)
//
// Returns the normalized name
func NormalizeName(name string, env int) (string, error) {
	// First remove any existing prefix/suffix
	re := regexp.MustCompile(`^(?:EPFL - )?(.*?)(?:\s*-\s*(?i:TEST|DEV))?$`)
	matches := re.FindStringSubmatch(name)

	// Use the captured group if there's a match, otherwise use the original name
	n := name
	if len(matches) > 1 {
		n = matches[1]
	}

	// Ensure bare name contains no " - " => replace with "-"
	re = regexp.MustCompile(` - `)
	n = re.ReplaceAllString(n, "-")

	switch env {
	case 1:
		return fmt.Sprintf("EPFL - %s - DEV", n), nil
	case 2:
		return fmt.Sprintf("EPFL - %s - TEST", n), nil
	case 3:
		return fmt.Sprintf("EPFL - %s", n), nil
	default:
		return "", fmt.Errorf("invalid environment: %d, must be 1, 2 or 3", env)
	}
}

// NormalizeURI performs some modification (required by Microsoft API) on URI
//   - removes the trailing slash from a string
//   - replace http with https
func NormalizeURI(s string) string {
	var n string
	if len(s) > 0 && s[len(s)-1] == '/' {
		n = s[:len(s)-1]
		s = n
	}

	if len(s) > 5 && s[:5] == "http:" {
		n = "https:" + s[5:]
		s = n
	}

	return s
}

// NormalizeThumbprint removes spaces and dashes from a thumbprint
func NormalizeThumbprint(thumbprint string) string {
	re, _ := regexp.Compile(`[\s\-]`)
	thumbprint = re.ReplaceAllString(thumbprint, "")

	return thumbprint
}
