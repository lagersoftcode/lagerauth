package api

import (
	"encoding/json"
	"io/ioutil"
	"lagerauth/logic"
	"lagerauth/models/core"
	"net/http"
	"strconv"
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

	numID, _ := strconv.Atoi(id)

	switch {
	case (path == "/" || path == "") && r.Method == http.MethodGet:
		listUsers(w, r)
	case id != "" && r.Method == http.MethodGet:
		viewUser(w, r, uint(numID))
	case id == "" && r.Method == http.MethodPost:
		createUser(w, r)
	case id != "" && r.Method == http.MethodPut:
		updateUser(w, r, uint(numID))
	case id != "" && r.Method == http.MethodDelete:
		deleteUser(w, r, uint(numID))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users := logic.GetUsers()
	json.NewEncoder(w).Encode(&users)
}

func viewUser(w http.ResponseWriter, r *http.Request, id uint) {
	user, err := logic.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload core.User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logic.CreateUser(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateUser(w http.ResponseWriter, r *http.Request, id uint) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload core.User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logic.UpdateUser(payload, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id uint) {
	err := logic.DeleteUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
