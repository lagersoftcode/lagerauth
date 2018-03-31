package permissions

import (
	"lagerauth/logic"
	"net/http"
)

type logoff struct{}

var Logoff *logoff

func (l *logoff) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	header := r.Header.Get("Authorization")
	token, err := logic.ValidateAndGetTokenFromAuthorizationHeader(header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = logic.DisableToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
