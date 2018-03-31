package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"lagerauth/logic"
	"lagerauth/bindata"
)

type resetPass struct{}

var ResetPass resetPass

func (rp *resetPass) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	pageData := passResetPage{}
	renderForgotPasswordPage(w, &pageData)
}

func (rp *resetPass) SendEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pageData := passResetPage{}

	err := logic.EmailExists(email)
	if err != nil {
		pageData.ErrorMessage = err.Error()
	} else {
		uuid := logic.AddPasswordResetCode(email)
		go logic.SendPasswordResetEmail(email, uuid)
		pageData.InfoMessage = fmt.Sprintf("Password reset mail sent to: %s", email)
		logic.Logger().Warn("Password reset mail sent to %s", email)
	}

	renderForgotPasswordPage(w, &pageData)
}

func (rp *resetPass) ResetPassword(w http.ResponseWriter, r *http.Request) {

	code := uuid.FromStringOrNil(r.URL.Query().Get("code"))

	pageData := passResetPage{}
	if code == uuid.Nil {
		pageData.DisableForm = true
		pageData.ErrorMessage = "Code is empty"
		renderResetPasswordPage(w, &pageData)
		return
	}

	if !logic.ResetPasswordCodeIsValid(code) {
		pageData.DisableForm = true
		pageData.ErrorMessage = "Code is not valid"
		renderResetPasswordPage(w, &pageData)
		return
	}

	pageData.Code = code.String()
	renderResetPasswordPage(w, &pageData)
}

func (rp *resetPass) ConfirmPassword(w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("password")
	passwordConfirm := r.FormValue("passwordConfirmation")
	code := uuid.FromStringOrNil(r.FormValue("code"))

	pageData := passResetPage{}

	if code == uuid.Nil {
		pageData.ErrorMessage = "Code not present, try again."
		pageData.DisableForm = true
		renderResetPasswordPage(w, &pageData)
		return
	}

	if password != passwordConfirm {
		pageData.ErrorMessage = "Passwords doesnt match"
		pageData.Code = code.String()
		renderResetPasswordPage(w, &pageData)
		return
	}

	err := logic.ResetPasswordFromCode(code, password)
	if err != nil {
		pageData.ErrorMessage = err.Error()
		pageData.DisableForm = true
		renderResetPasswordPage(w, &pageData)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: oAuthLoginCookieName, MaxAge: 0, Expires: time.Now().Add(-100 * time.Hour)})
	pageData.InfoMessage = "Password changed sucessfully"
	pageData.DisableForm = true
	renderResetPasswordPage(w, &pageData)
}

type passResetPage struct {
	InfoMessage  string
	ErrorMessage string
	DisableForm  bool
	Title        string
	Code         string
}

func renderResetPasswordPage(w http.ResponseWriter, pageData *passResetPage) {
	resetPassView, _ := bindata.Asset("views/resetpass.html.tmpl")
	layoutView, _ := bindata.Asset("views/layout.html.tmpl")
	
	t, err := template.New("oAuth").Parse(string(resetPassView) + string(layoutView))
	if err != nil {
		log.Panic(err)
	}
	pageData.Title = "Reset Password"
	t.ExecuteTemplate(w, "layout", &pageData)
}

func renderForgotPasswordPage(w http.ResponseWriter, pageData *passResetPage) {
	
	forgotView, _ := bindata.Asset("views/forgotpass.html.tmpl")
	layoutView, _ := bindata.Asset("views/layout.html.tmpl")
	
	t, err := template.New("oAuth").Parse(string(forgotView) + string(layoutView))
	if err != nil {
		log.Panic(err)
	}
	pageData.Title = "Forgot Password"
	t.ExecuteTemplate(w, "layout", &pageData)
}
