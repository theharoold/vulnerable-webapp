package main

import (
	"bzrm/session-hijacking/api"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	r := http.NewServeMux()

	var a *api.API = new(api.API)
	a.SessionTokens = make([]string, 0, 10)

	r.HandleFunc("/", a.HomepageHandler)
	r.HandleFunc("/login", a.LoginHandler)
	r.HandleFunc("/profile", a.ProfileHandler)

	http.ListenAndServe(":5000", r)
}
