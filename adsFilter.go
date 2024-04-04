/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Decide which Ad to response
 * 				according to user's data.
 *
 */

package main

import (
	"encoding/json"

	"github.com/Ateto1204/Ads-Delivery-Service/models"
)

func filterAds(ads []models.Ad, queryParams map[string][]string) []models.ResponseAd {
	var filteredAds []models.ResponseAd

	for _, ad := range ads {
		if matchesQuery(ad.Conditions, queryParams) {
			responseAd := models.ResponseAd{ad.Title, ad.EndAt}
			filteredAds = append(filteredAds, responseAd)
		}
	}

	return filteredAds
}

func matchesQuery(conditions models.Conditions, queryParams map[string][]string) bool {
	if ageParam, ok := queryParams["age"]; ok {
		var age int
		if err := json.Unmarshal([]byte(ageParam[0]), &age); err != nil {
			return false
		}
		if conditions.Age.Start != 0 && (age < conditions.Age.Start || age > conditions.Age.End) {
			return false
		}
	}

	if genderParam, ok := queryParams["gender"]; ok && conditions.Gender != "" && genderParam[0] != conditions.Gender {
		return false
	}

	if countryParam, ok := queryParams["country"]; ok && len(conditions.Country) > 0 {
		found := false
		for _, country := range conditions.Country {
			if contains(countryParam, country) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if platformParam, ok := queryParams["platform"]; ok && len(conditions.Platform) > 0 {
		found := false
		for _, platform := range conditions.Platform {
			if contains(platformParam, platform) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func contains(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}
