package api

import (
	"fmt"
	"net/http"
)

func (a *API) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Get the username and password from the request
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Invalid basic authentication credentials", http.StatusBadRequest)
			return
		}

		// Validate the login credentials
		if username == "" || password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}
		if !Authenticate(username, password) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Generate a session token and set it as a cookie in the response
		token := generateSessionToken()

		// Avoid nil dereferencing
		if a.SessionTokens == nil {
			a.SessionTokens = make([]string, 0, 10)
		}

		// Add session token to the list
		a.SessionTokens = append(a.SessionTokens, token)
		fmt.Println(a.SessionTokens)

		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: token,
			Path:  "/",
		})

		// Return a success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Login successful")
	}
}
