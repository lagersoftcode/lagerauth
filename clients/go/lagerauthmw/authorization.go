package lagerauthmw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Authorize handles the authorization part of our api its an http middleware
func (l LagerAuthMiddleware) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if l.can(r) {
			next.ServeHTTP(w, r)
		} else {
			l.unauthorized(w, r)
		}
	})
}

type can struct {
	Method     string
	Controller string
	Action     string
	Extra      string
}

func (l LagerAuthMiddleware) can(r *http.Request) bool {

	token, err := getJWT(r, l.config)
	if err != nil {
		return false
	}

	// setup can model
	method := r.Method

	controller := "home"
	action := "index"

	if containsControllerAndAction(r.URL.Path) {
		parts := strings.Split(r.URL.Path, "/")

		if isNotAFile(parts[1]) {
			controller = parts[1]
		}

		if isNotAFile(parts[2]) {
			action = removeQueryPart(parts[2])
		}
	} else if containsControllerOnly(r.URL.Path) {
		parts := strings.Split(r.URL.Path, "/")

		if isNotAFile(parts[1]) {
			controller = removeQueryPart(parts[1])
		}
	}

	json, _ := json.Marshal(&can{
		Method:     method,
		Controller: controller,
		Action:     action,
	})

	jsonStr := string(json)
	log.Println(jsonStr)

	// set-up request
	url := fmt.Sprintf("%s/can", l.OAuthURL)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(json))
	header := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", header)
	req.Header.Set("Content-Type", "application/json")

	// execute
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false
	}

	// validate if its a 200OK (authorized)
	return res.StatusCode == http.StatusOK
}

// redirect to /{mountPoint}/unauthorized
func (l LagerAuthMiddleware) unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, l.config.MountPoint+"/unauthorized", http.StatusTemporaryRedirect)
}

// helper functions for parsing URL.Path
func isNotAFile(s string) bool {
	return s != "" && !strings.Contains(s, ".")
}

func removeQueryPart(s string) string {
	if strings.Contains(s, "?") {
		s = strings.Split(s, "?")[0]
	}

	return s
}

func containsControllerAndAction(s string) bool {
	parts := strings.Split(s, "/")

	return len(parts) > 2
}

func containsControllerOnly(s string) bool {
	parts := strings.Split(s, "/")

	return len(parts) > 1
}
