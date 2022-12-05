package api

import (
	"fmt"
	"net/http"
)

func (a *API) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check if the user is logged in by looking for the session token cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "You are not logged in", http.StatusUnauthorized)
		return
	}

	// Validate the session token
	if !validateSessionToken(cookie.Value, a.SessionTokens) {
		http.Error(w, "Invalid session token", http.StatusUnauthorized)
		return
	}

	// Return the profile information
	fmt.Fprintln(w, "{ 'data': 'This is some very private profile data' }")
}
