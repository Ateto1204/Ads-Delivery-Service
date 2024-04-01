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
	"fmt"
	"net/http"
	"sort"
	"time"
)

func listAdsHandler(w http.ResponseWriter, r *http.Request) {
	var activeAds []Ad
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

	response := map[string][]ResponseAd{"items": filteredAds}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("GET FAIL")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getPaginationParams(r *http.Request) (int, int) {
	offset := 1
	limit := 5

	if offsetParam := r.URL.Query().Get("offset"); offsetParam != "" {
		fmt.Sscanf(offsetParam, "%d", &offset)
	}

	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		fmt.Sscanf(limitParam, "%d", &limit)
	}

	// Validate offset and limit
	if offset < 1 {
		offset = 1
	}
	if limit < 1 || limit > 100 {
		limit = 5
	}

	return offset, limit
}

func calculatePaginationRange(offset, limit, total int) (int, int) {
	start := (offset - 1) * limit
	end := start + limit
	if end > total {
		end = total
	}
	return 0, end
}
