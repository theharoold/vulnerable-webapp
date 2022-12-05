package api

import "fmt"

func validateSessionToken(token string, sessionTokens []string) bool {
	// Look up the session token in a database or cache
	// If the token is found, it is considered valid
	return contains(token, sessionTokens)
}

// Iterate over the slice and check if it contains
// the given string
func contains(s string, slice []string) bool {
	for _, value := range slice {
		if s == value {
			fmt.Println("s: " + s + ", value: " + value)
			return true
		}
	}

	return false
}
