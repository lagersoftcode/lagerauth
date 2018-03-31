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

type applications struct{}

func NewApplications() *applications {
	return &applications{}
}

func (applications *applications) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// dispatch methods depending on the routes
	path := strings.TrimPrefix(r.URL.Path, "/applications")
	pathSplitted := strings.Split(path, "/")

	id := pathSplitted[len(pathSplitted)-1]
	numID, _ := strconv.Atoi(id)

	switch {
	case (path == "/" || path == "") && r.Method == http.MethodGet:
		listApplications(w, r)
	case id != "" && r.Method == http.MethodGet:
		viewApplication(w, r, uint(numID))
	case id == "" && r.Method == http.MethodPost:
		createApplication(w, r)
	case id != "" && r.Method == http.MethodPut:
		updateApplication(w, r, uint(numID))
	case id != "" && r.Method == http.MethodDelete:
		deleteApplication(w, r, uint(numID))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func listApplications(w http.ResponseWriter, r *http.Request) {
	apps, err := logic.GetApplications()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&apps)
}

func viewApplication(w http.ResponseWriter, r *http.Request, id uint) {
	app, err := logic.GetApplication(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&app)
}

func createApplication(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload core.Application
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logic.CreateApplication(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateApplication(w http.ResponseWriter, r *http.Request, id uint) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var payload core.Application
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logic.UpdateApplication(payload, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteApplication(w http.ResponseWriter, r *http.Request, id uint) {
	err := logic.DeleteApplication(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
