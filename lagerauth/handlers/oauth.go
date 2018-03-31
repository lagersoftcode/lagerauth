package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"net/url"

	"github.com/satori/go.uuid"

	"lagerauth/bindata"
	"lagerauth/logic"
	"lagerauth/models"

	"github.com/dgrijalva/jwt-go"
)

var oAuthLoginCookieName = "oAuthEmail"

type oauth struct{}

var OAuth *oauth

/*AuthHandler job is to render the login form, if the user is not logged in.
 * if the user is logged in redirect with code (skipping authorization of use account)
 */
func (o *oauth) AuthHandler(w http.ResponseWriter, r *http.Request) {

	redirectURI := r.URL.Query().Get("redirect_uri")
	clientID, _ := uuid.FromString(r.URL.Query().Get("client_id"))
	state := r.URL.Query().Get("state")

	redirectURI = addStateToUri(state, redirectURI)
	end := handleMissingOAuthParams(w, redirectURI, clientID)
	if end {
		return
	}

	authPage := &authPage{RedirectURI: redirectURI, ClientID: clientID.String()}
	oAuthLoginCookie, err := r.Cookie(oAuthLoginCookieName)
	if err != nil {
		renderOAuthPage(w, authPage)
		return
	}

	token, err := jwt.Parse(oAuthLoginCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("Invalid Signature Method")
		}

		return logic.JWTSecretKey(), nil
	})

	if !token.Valid {
		authPage.ErrorMessage = "JWT validation failed"
		renderOAuthPage(w, authPage)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email := claims["email"].(string)
		getCodeAndRedirect(w, r, email, redirectURI, clientID)
	}

	renderOAuthPage(w, authPage)
}

/*LoginHandler handles the login form post of AuthHanlder */
func (o *oauth) LoginHandler(w http.ResponseWriter, r *http.Request) {

	redirectURI := r.URL.Query().Get("redirect_uri")
	clientID, _ := uuid.FromString(r.URL.Query().Get("client_id"))
	state := r.URL.Query().Get("state")

	redirectURI = addStateToUri(state, redirectURI)
	end := handleMissingOAuthParams(w, redirectURI, clientID)
	if end {
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	err := logic.LoginValid(email, password)
	if err != nil {
		renderOAuthPage(w, &authPage{RedirectURI: redirectURI, ClientID: clientID.String(), ErrorMessage: err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email})
	tokenString, err := token.SignedString(logic.JWTSecretKey())

	http.SetCookie(w, &http.Cookie{Name: oAuthLoginCookieName, Value: tokenString}) // add logged in cookie
	getCodeAndRedirect(w, r, email, redirectURI, clientID)
}

type token struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
}

/*TokenHandler handles the exchange of the code, for a valid token */
func (o *oauth) TokenHandler(w http.ResponseWriter, r *http.Request) {
	code, _ := uuid.FromString(r.FormValue("code"))
	clientID, _ := uuid.FromString(r.FormValue("client_id"))
	clientSecret := r.FormValue("client_secret")

	var tokenStr string
	var err error
	if clientSecret == "" {
		tokenStr, err = logic.GenerateTokenNoSecret(code, clientID)
	} else {
		tokenStr, err = logic.GenerateToken(code, clientID, clientSecret)
	}

	encoder := json.NewEncoder(w)

	if err != nil {
		message := models.NewError(err)
		encoder.Encode(message)
		return
	}

	token := token{Token: tokenStr, Type: "Bearer"} // Always type bearer
	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(token)
}

type authPage struct {
	RedirectURI  string
	ClientID     string
	ErrorMessage string
	InfoMessage  string
	DisableLogin bool
	Title        string
}

func renderOAuthPage(w http.ResponseWriter, pageData *authPage) {
	authView, _ := bindata.Asset("views/auth.html.tmpl")
	layoutView, _ := bindata.Asset("views/layout.html.tmpl")

	t, _ := template.New("oAuth").Parse(string(authView) + string(layoutView))
	pageData.Title = "Login"
	t.ExecuteTemplate(w, "layout", &pageData)
}

func handleMissingOAuthParams(w http.ResponseWriter, redirectURI string, clientID uuid.UUID) bool {
	if redirectURI == "" || clientID == uuid.Nil {
		authPage := &authPage{RedirectURI: redirectURI, ClientID: clientID.String()}
		authPage.ErrorMessage = "client_id or redirect_uri is blank, this should never happen. Please go back to the original application and try to login again."
		authPage.DisableLogin = true
		renderOAuthPage(w, authPage)
		return true
	}

	return false
}

func getCodeAndRedirect(w http.ResponseWriter, r *http.Request, email, redirectURI string, clientID uuid.UUID) {
	url, _ := url.Parse(redirectURI)
	q := url.Query()

	code, err := logic.GenerateCode(email, clientID)
	if err != nil {
		renderOAuthPage(w, &authPage{ErrorMessage: err.Error(), DisableLogin: true})
		return
	}

	q.Add("code", code.String())
	url.RawQuery = q.Encode()
	r.Method = http.MethodGet
	http.Redirect(w, r, url.String(), http.StatusSeeOther)
}

func addStateToUri(state, redirectURI string) string {
	if state == "" {
		return redirectURI
	}

	uri, _ := url.ParseRequestURI(redirectURI)
	q := uri.Query()
	q.Add("state", state)
	uri.RawQuery = q.Encode()
	return uri.String()
}
