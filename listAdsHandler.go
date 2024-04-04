/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Accept API request with GET,
 *				and reponse the Ads.
 *
 */

package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/Ateto1204/Ads-Delivery-Service/models"
)

func listAdsHandler(w http.ResponseWriter, r *http.Request) {
	var activeAds []models.Ad
	now := time.Now()

	mu.Lock()
	defer mu.Unlock()

	for _, ad := range Ads {
		if ad.StartAt.Before(now) && ad.EndAt.After(now) {
			activeAds = append(activeAds, ad)
		}
	}

	// Sorting by EndAt in ascending order
	if len(activeAds) > 1 {
		sort.Slice(activeAds, func(i, j int) bool {
			return activeAds[i].EndAt.Before(activeAds[j].EndAt)
		})
	}

	// Pagination
	offset, limit := getPaginationParams(r)
	start, end := calculatePaginationRange(offset, limit, len(activeAds))

	// Filtering based on query parameters
	filteredAds := filterAds(activeAds[start:end], r.URL.Query())

	response := map[string][]models.ResponseAd{"items": filteredAds}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
