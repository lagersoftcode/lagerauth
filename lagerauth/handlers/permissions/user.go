package permissions

import (
	"encoding/json"
	"lagerauth/logic"
	"net/http"
)

type user struct{}

var User *user = &user{}

func (u *user) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	token, err := logic.ValidateAndGetTokenFromAuthorizationHeader(header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := logic.GetUserFromToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	info := omniauthHash{
		Provider: "lagerauth",
		UID:      user.Email,
		Info: userInfo{
			Name:  user.Email,
			Email: user.Email,
		},
	}

	json.NewEncoder(w).Encode(info)
}

type omniauthHash struct {
	Provider string   `json:"provider"`
	UID      string   `json:"uid"`
	Info     userInfo `json:"info"`
}

type userInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
