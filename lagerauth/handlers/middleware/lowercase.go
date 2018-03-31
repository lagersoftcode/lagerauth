package middleware

import (
	"net/http"
	"strings"
)

/*Lowercase makes all of the routes lowercase internally, /route != /Route in go (or the html standard) but in most languages they are */
func Lowercase(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
