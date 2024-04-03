/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Accept API request with POST,
 * 				and store the Ad.
 *
 */

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func createAdHandler(w http.ResponseWriter, r *http.Request) {
	var ad Ad
	err := json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	Ads = append(Ads, ad)

	w.WriteHeader(http.StatusCreated)

	log.Println(ad)
}
