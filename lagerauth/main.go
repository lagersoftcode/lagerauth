package main

import (
	"log"
	"net/http"
	"os"

	"lagerauth/config"
	"lagerauth/database"
	"lagerauth/email"
	"lagerauth/handlers"
	"lagerauth/handlers/api"
	"lagerauth/handlers/middleware"
	"lagerauth/handlers/permissions"
	"lagerauth/logger"
	"lagerauth/logic"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.Read()          //read configuration
	db := database.New(conf)       //connect to db
	emailSender := email.New(conf) //email sender

	logger := logger.New(logger.Options{Writer: os.Stdout, Database: db.DB.DB()})

	logic.SetLogger(logger)           //set package-level 'logger' variable on logic package
	logic.SetDB(db)                   //set package-level 'db' variable on logic package
	logic.SetJWTSecret(conf.JWTKey)   //set package-level 'jwtKey' variable on logic package
	logic.SetEmailSender(emailSender) //set package-level 'emailSender' variable on logic package

	r := mux.NewRouter()

	/* OAuth 2.0 Handlers */
	r.HandleFunc("/auth", handlers.OAuth.AuthHandler).Methods(http.MethodGet)
	r.HandleFunc("/auth", handlers.OAuth.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/token", handlers.OAuth.TokenHandler).Methods(http.MethodPost)

	/* Password Reset */
	r.HandleFunc("/forgotpass", handlers.ResetPass.ForgotPassword).Methods(http.MethodGet)
	r.HandleFunc("/forgotpass", handlers.ResetPass.SendEmail).Methods(http.MethodPost)
	r.HandleFunc("/resetpass", handlers.ResetPass.ResetPassword).Methods(http.MethodGet)
	r.HandleFunc("/resetpass", handlers.ResetPass.ConfirmPassword).Methods(http.MethodPost)

	/* Permission Server Routes
	 * Can: returns 200 or 401/403 depending if the user has access to the resource
	 * Logoff: returns 200 always disables the user token
	 * User: returns a omniauth user info hash (for gitlab)
	 * Menu: returns a list of resources user has access in that application so that we can render menus more efficiently
	 */
	r.Handle("/can", middleware.WithAuthentication(middleware.JSONContentType(permissions.Can))).Methods(http.MethodPost) //we manage authorization on the route.
	r.Handle("/logoff", middleware.WithAuthentication(middleware.JSONContentType(permissions.Logoff))).Methods(http.MethodPost)
	r.Handle("/user", middleware.WithAuthentication(middleware.JSONContentType(permissions.User))).Methods(http.MethodPost, http.MethodGet)
	r.Handle("/menu", middleware.WithAuthentication(middleware.JSONContentType(permissions.Menu))).Methods(http.MethodPost)

	/* Api Routes For OAuth Site: */
	apiRouter := mux.NewRouter()
	r.PathPrefix("/api").Handler(http.StripPrefix("/api", middleware.WithAuthorization(middleware.JSONContentType(apiRouter)))) //we manage authorization on the middleware.
	apiRouter.PathPrefix("/applications").Handler(api.NewApplications())
	apiRouter.PathPrefix("/users").Handler(api.NewUsers())

	/* Static files for oauth*/
	r.PathPrefix("/assets").Handler(handlers.NewAssetsHandler())

	/* Static files and setup for vue.js */
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("wwwroot/static"))))
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wwwroot/index.html") //this handles 404's needed for vue-router history mode. (which in turn is needed for the redirect_uri to work correctly.)
	})

	log.Printf("Starting lagerauth... bind: %s\n", conf.Bind)
	err := http.ListenAndServe(conf.Bind, middleware.Lowercase(middleware.CORS(r)))
	if err != nil {
		log.Fatal(err.Error())
	}
}
