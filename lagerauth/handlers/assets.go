package handlers

import (
	"strings"
	"net/http"

	"lagerauth/bindata"
)

type AssetsHandler struct{}

func NewAssetsHandler() *AssetsHandler {
	return &AssetsHandler{}
}

func (a *AssetsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := strings.TrimPrefix(r.URL.Path, "/")
	file, err := bindata.Asset(request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(file)
}