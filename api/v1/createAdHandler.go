/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Accept API request with POST,
 * 				and store the Ad.
 *
 */

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Ateto1204/Ads-Delivery-Service/models"
)

func createAdHandler(w http.ResponseWriter, r *http.Request) {
	var ad models.Ad
	err := json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo.muxLock()
	repo.muxUnlock()

	repo.addAd(ad)

	w.WriteHeader(http.StatusCreated)
}
