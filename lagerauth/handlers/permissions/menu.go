package permissions

import "net/http"

type menu struct{}

var Menu *menu = &menu{}

func (m *menu) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// return a list of the resources
	w.WriteHeader(http.StatusOK)
}
