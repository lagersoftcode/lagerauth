package api

import (
	"encoding/json"
	"io/ioutil"
	"lagerauth/logic"
	"net/http"
	"strconv"
	"strings"

	"lagerauth/models/core"
)

type roles struct{}

func NewRoles() *roles {
	return &roles{}
}

func (roles *roles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// dispatch methods depending on the routes
	path := strings.TrimPrefix(r.URL.Path, "/roles")
	pathSplitted := strings.Split(path, "/")

	id := pathSplitted[len(pathSplitted)-1]
	numID, _ := strconv.Atoi(id)

	switch {
	case (path == "/" || path == "") && r.Method == http.MethodGet:
		listRoles(w, r)
	case id != "" && r.Method == http.MethodGet:
		viewRole(w, r, uint(numID))
	case id == "" && r.Method == http.MethodPost:
		createRole(w, r)
	case id != "" && r.Method == http.MethodPut:
		updateRole(w, r, uint(numID))
	case id != "" && r.Method == http.MethodDelete:
		deleteRole(w, r, uint(numID))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func listRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := logic.GetRoles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&roles)
}

func viewRole(w http.ResponseWriter, r *http.Request, id uint) {
	role, err := logic.GetRole(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&role)
}

func createRole(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload core.Role
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logic.CreateRole(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRole(w http.ResponseWriter, r *http.Request, id uint) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload core.Role
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logic.UpdateRole(payload, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRole(w http.ResponseWriter, r *http.Request, id uint) {
	err := logic.DeleteRole(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
