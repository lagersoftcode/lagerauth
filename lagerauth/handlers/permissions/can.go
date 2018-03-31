package permissions

import (
	"encoding/json"
	"io/ioutil"
	"lagerauth/logic"
	"net/http"
	"strings"
)

var Can *can = &can{}

type can struct {
	Method     string
	Controller string
	Action     string
	Extra      string
}

func (c *can) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this should only return 200OK when user has access and 403 Forbidden when user has no access.
	// middleware handles authentication with a 401 Unathorized but we check again for good measure.
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload can
	err = json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	header := r.Header.Get("Authorization")
	token, err := logic.ValidateAndGetTokenFromAuthorizationHeader(header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check for permissions
	permissions, err := logic.GetApplicationPermissionsFromToken(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var valid bool
	for _, v := range permissions {
		valid = valid || ((payload.Controller == strings.ToLower(v.Controller) || v.Controller == "*") &&
			(payload.Action == strings.ToLower(v.Action) || v.Action == "*") &&
			(payload.Method == strings.ToLower(v.Method) || v.Method == "*"))
	}

	if valid {
		w.WriteHeader(http.StatusOK)
		return
	}

	// no permission:
	w.WriteHeader(http.StatusForbidden)
}
