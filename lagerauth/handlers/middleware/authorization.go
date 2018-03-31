package middleware

import (
	"lagerauth/logic"

	"net/http"
)

/*WithAuthorization checks authentication and authorization to a resource*/
func WithAuthorization(next http.Handler) http.Handler {
	return WithAuthentication(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			header := r.Header.Get("Authorization")
			token, err := logic.ValidateAndGetTokenFromAuthorizationHeader(header)
			get user permissions with token
			check if has permissions
			Fail with http.StatusForbidden if user doesnt have permissions
			else serve next
		*/
		next.ServeHTTP(w, r)
	}))
}

/*WithAuthentication checks only authorization */
func WithAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		_, err := logic.ValidateAndGetTokenFromAuthorizationHeader(header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
