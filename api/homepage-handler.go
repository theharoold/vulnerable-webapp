package api

import (
	"fmt"
	"net/http"
)

func (a *API) HomepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "text/html")

	// Return the profile information
	fmt.Fprintln(w, "<h1>BZRM example</h1>")
}
