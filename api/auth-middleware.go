package api

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// Authenticate validates the given username and password using Basic Auth
func Authenticate(username, password string) bool {
	// TODO: Add code to validate the given username and password against a database or other store

	// For now, just return true if the username is "veljko" or "savo" and the password is "password"
	return (username == "veljko" || username == "savo") && password == "password"
}

// BasicAuthMiddleware is a middleware that checks the HTTP Authorization header for a valid Basic Auth
// username and password. If the header is present and the username and password are valid, the request
// is allowed to continue. Otherwise, a 401 Unauthorized response is returned.
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Basic Auth credentials from the request header
		auth := r.Header.Get("Authorization")

		// If the header is not present, return a 401 Unauthorized response
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// The header is present, so get the username and password from it
		credentials, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		if err != nil {
			// If there was an error decoding the credentials, return a 401 Unauthorized response
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Split the username and password at the colon (:) character
		parts := strings.SplitN(string(credentials), ":", 2)
		if len(parts) != 2 {
			// If the credentials are not in the correct format, return a 401 Unauthorized response
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Get the username and password from the split credentials
		username, password := parts[0], parts[1]

		// Check if the username and password are valid
		if !Authenticate(username, password) {
			// If the credentials are invalid, return a 401 Unauthorized response
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// If we made it here, the credentials are valid, so call the next handler
		next.ServeHTTP(w, r)
	})
}
