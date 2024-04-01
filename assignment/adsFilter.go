/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Decide which Ad to response
 * 				according to user's data.
 *
 */

package main

import "encoding/json"

func filterAds(ads []Ad, queryParams map[string][]string) []ResponseAd {
	var filteredAds []ResponseAd

	for _, ad := range ads {
		if matchesQuery(ad.Conditions, queryParams) {
			responseAd := ResponseAd{ad.Title, ad.EndAt}
			filteredAds = append(filteredAds, responseAd)
		}
	}

	return filteredAds
}

func matchesQuery(conditions Conditions, queryParams map[string][]string) bool {
	// Implement logic to check if conditions match queryParams
	// For simplicity, we are assuming that all conditions are AND-ed.

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
