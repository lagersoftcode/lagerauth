package api

import (
	"encoding/json"
	"lagerauth/logic"
	"net/http"
	"strings"
)

type users struct{}

func NewUsers() *users {
	return &users{}
}

func (users *users) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// dispatch methods depending on the routes
	path := strings.TrimPrefix(r.URL.Path, "/users")
	pathSplitted := strings.Split(path, "/")
	id := pathSplitted[len(pathSplitted)-1]

	switch {
	case (path == "/" || path == "") && r.Method == http.MethodGet:
		listUsers(w, r)
	case id != "" && r.Method == http.MethodGet:
		viewUser(w, r)
	case id == "" && r.Method == http.MethodPost:
		createUser(w, r)
	case id != "" && r.Method == http.MethodPut:
		updateUser(w, r)
	case id != "" && r.Method == http.MethodDelete:
		deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users := logic.GetUsers()
	json.NewEncoder(w).Encode(&users)
}

func viewUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Api View"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Api Create"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Api Update"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Api Delete"))
}
