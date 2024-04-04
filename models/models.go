/*
 * Project: 2024 Dcard Backend Intern Assignment
 * Author: Ateto
 *
 * Description: Define the structure of Ad.
 *
 */

package models

import "time"

type Ad struct {
	Title      string     `json:"title"`
	StartAt    time.Time  `json:"startAt"`
	EndAt      time.Time  `json:"endAt"`
	Conditions Conditions `json:"conditions"`
}

type ResponseAd struct {
	Title string    `json:"title"`
	EndAt time.Time `json:"endAt"`
}

type Conditions struct {
	Age      AgeRange `json:"age,omitempty"`
	Gender   string   `json:"gender,omitempty"`
	Country  []string `json:"country,omitempty"`
	Platform []string `json:"platform,omitempty"`
}

type AgeRange struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}
