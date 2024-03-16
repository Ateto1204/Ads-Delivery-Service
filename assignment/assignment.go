/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Recommend advertisements to customers
 * 				according to some information of the
 * 				customers.
 *
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Ad represents an advertisement.
type Ad struct {
	Title      string     `json:"title"`
	StartAt    time.Time  `json:"startAt"`
	EndAt      time.Time  `json:"endAt"`
	Conditions Conditions `json:"conditions"`
}

// Conditions represent the conditions for displaying an advertisement.
type Conditions struct {
	// Age      AgeRange `json:"age,omitempty"`
	Gender   string   `json:"gender,omitempty"`
	AgeStart int      `json:"ageStart"`
	AgeEnd   int      `json:"ageEnd"`
	Country  []string `json:"country,omitempty"`
	Platform []string `json:"platform,omitempty"`
}

// AgeRange represents the age range for the conditions.
type AgeRange struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// Ads is an in-memory store for advertisements.
var Ads []Ad
var mu sync.Mutex

func createAdHandler(w http.ResponseWriter, r *http.Request) {
	var ad Ad
	err := json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		fmt.Println("------------------------------\nPOST FAIL\n-----------------------")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	Ads = append(Ads, ad)

	w.WriteHeader(http.StatusCreated)

	fmt.Println("------------------------------\nPOST SUCCESS\n-----------------------")
	fmt.Println("Received ad:", ad)
}

func listAdsHandler(w http.ResponseWriter, r *http.Request) {
	var activeAds []Ad
	// now := time.Now()

	mu.Lock()
	defer mu.Unlock()

	for _, ad := range Ads {
		// if ad.StartAt.Before(now) && ad.EndAt.After(now) {
		activeAds = append(activeAds, ad)
		// }
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
	fmt.Println("FILTER SIZE: ", len(filteredAds))

	response := map[string][]Ad{"items": filteredAds}
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
	return 1, 4
}

func filterAds(ads []Ad, queryParams map[string][]string) []Ad {
	var filteredAds []Ad

	for _, ad := range ads {
		if matchesQuery(ad.Conditions, queryParams) {
			filteredAds = append(filteredAds, ad)
		}
	}

	return filteredAds
}

func matchesQuery(conditions Conditions, queryParams map[string][]string) bool {
	// Implement logic to check if conditions match queryParams
	// For simplicity, we are assuming that all conditions are AND-ed.

	// if ageParam, ok := queryParams["age"]; ok {
	// 	var age AgeRange
	// 	if err := json.Unmarshal([]byte(ageParam[0]), &age); err != nil {
	// 		return false
	// 	}
	// 	if conditions.Age.Start != 0 && (age.Start < conditions.Age.Start || age.End > conditions.Age.End) {
	// 		return false
	// 	}
	// }

	// if genderParam, ok := queryParams["gender"]; ok && conditions.Gender != "" && genderParam[0] != conditions.Gender {
	// 	return false
	// }

	// if countryParam, ok := queryParams["country"]; ok && len(conditions.Country) > 0 {
	// 	for _, c := range conditions.Country {
	// 		if !contains(countryParam, c) {
	// 			return false
	// 		}
	// 	}
	// }

	// if platformParam, ok := queryParams["platform"]; ok && len(conditions.Platform) > 0 {
	// 	for _, p := range conditions.Platform {
	// 		if !contains(platformParam, p) {
	// 			return false
	// 		}
	// 	}
	// }

	return true
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ad", createAdHandler).Methods("POST")
	r.HandleFunc("/api/v1/ad", listAdsHandler).Methods("GET")

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
