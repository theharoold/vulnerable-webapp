package api

import (
	"math/rand"
)

// generateSessionToken generates a new session token
func generateSessionToken() string {
	return generateRandomString(32)
}

func generateRandomString(length int) string {
	// Initialize a string of all possible alphanumeric characters
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Initialize a byte slice of the specified length
	b := make([]byte, length)

	// Loop through the byte slice and set each element to a random character from the string of all possible characters
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}

	// Return the generated random string
	return string(b)
}
