package api

import (
	"net/http"
)

// AuthMiddleware is a middleware that checks if the user is authenticated.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement authentication
		next.ServeHTTP(w, r)
	})
}
