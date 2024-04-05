package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setUpRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ad", createAdHandler).Methods("POST")
	r.HandleFunc("/api/v1/ad", listAdsHandler).Methods("GET")

	return r
}
