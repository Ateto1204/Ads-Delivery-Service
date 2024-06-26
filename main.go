/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Recommend advertisements to customers
 * 				according to some information of the
 * 				customers.
 *
 * go build *.go
 */

package main

import (
	"net/http"
	"sync"

	"github.com/Ateto1204/Ads-Delivery-Service/models"
	"github.com/gorilla/mux"
)

// Ads is an in-memory store for advertisements.
var Ads []models.Ad
var mu sync.Mutex

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ad", createAdHandler).Methods("POST")
	r.HandleFunc("/api/v1/ad", listAdsHandler).Methods("GET")

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
