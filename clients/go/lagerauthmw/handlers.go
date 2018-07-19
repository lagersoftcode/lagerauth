package lagerauthmw

import (
	"net/http"
)

// RegisterLagerauthHandler takes a *mux and registers the lagerauth routes to make it work (as configured on mount point)
func (l LagerAuthMiddleware) RegisterLagerauthHandler(mux *http.ServeMux) {
	mux.Handle(l.config.MountPoint+"/login", loginHandler(l))
	mux.Handle(l.config.MountPoint+"/logout", logoutHandler(l))
	mux.Handle(l.config.MountPoint+"/auth", authHandler(l))
	mux.Handle(l.config.MountPoint+"/unauthorized", unauthorizedHandler(l))
}

func loginHandler(lmw LagerAuthMiddleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		scheme := "http"
		if r.URL.Scheme != "" {
			scheme = r.URL.Scheme
		}

		host := scheme + "://" + r.Host
		http.Redirect(w, r, lmw.LoginLink(host), http.StatusTemporaryRedirect)
	})
}

func logoutHandler(lmw LagerAuthMiddleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lmw.deAuthenticate(w, r)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}

func authHandler(lmw LagerAuthMiddleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")

		err := lmw.authenticate(code, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}

// TODO: add link to login
func unauthorizedHandler(lmw LagerAuthMiddleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
