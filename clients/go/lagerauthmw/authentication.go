package lagerauthmw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// LoginLink returns the link to lagerauth.
func (l LagerAuthMiddleware) LoginLink(host string) string {
	returnURL := fmt.Sprintf("%s%s/auth", host, l.config.MountPoint)
	encoded, _ := url.Parse(returnURL)

	return fmt.Sprintf("%s/auth?client_id=%s&redirect_uri=%s", l.config.OAuthURL, l.config.ClientID, encoded.String())
}

// token type for authentication
type token struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
}

// exchange code from token
// set cookie jwt with secretkey
func (l LagerAuthMiddleware) authenticate(code string, w http.ResponseWriter, r *http.Request) error {
	tokenBytes, err := getToken(code, l.config)
	if err != nil {
		return err
	}

	var token token
	err = json.Unmarshal(tokenBytes, &token)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    l.config.CookieName,
		Value:   token.Token,
		Expires: time.Now().AddDate(1, 0, 0),
		Path:    "/",
	})

	return nil
}

// invalidate session and remove cookie
func (l LagerAuthMiddleware) deAuthenticate(w http.ResponseWriter, r *http.Request) error {
	token, err := getJWT(r, l.config)

	if err == nil {
		url := fmt.Sprintf("%s/logoff", l.OAuthURL)
		req, _ := http.NewRequest(http.MethodPost, url, nil)
		header := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", header)

		client := &http.Client{}
		client.Do(req)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    l.config.CookieName,
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})

	return nil
}

// exchange code for token
func getToken(code string, conf config) ([]byte, error) {
	url := fmt.Sprintf("%s/token?client_id=%s&client_secret=%s&code=%s", conf.OAuthURL, conf.ClientID, conf.SecretKey, code)

	res, err := http.Post(url, "text/html", nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return body, err
}

// get JWT from cookies
func getJWT(r *http.Request, conf config) (string, error) {
	cookie, err := r.Cookie(conf.CookieName)

	if err != nil {
		return "", err
	}

	return cookie.Value, err
}
