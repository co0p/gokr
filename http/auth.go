package http

import (
	"crypto/subtle"
	"log"
	"net/http"
)

// WithBasicAuth wraps a given handlerFunc with basic auth protection
// copied shamelessly from https://stackoverflow.com/questions/21936332/idiomatic-way-of-requiring-http-basic-auth-in-go
func WithBasicAuth(handler http.HandlerFunc, username, password string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="gokrs"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			log.Printf("rejected access for user=%s pwd=%s\n", user, pass)
			return
		}

		log.Printf("granting access for user=%s pwd=%s\n", user, pass)

		handler(w, r)
	}
}
