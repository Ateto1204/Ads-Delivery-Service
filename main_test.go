package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Ateto1204/Ads-Delivery-Service/models"
)

func TestCreateAdHandler(t *testing.T) {
	ad := models.Ad{
		Title:   "Test Ad",
		StartAt: time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC),
		EndAt:   time.Date(2024, time.April, 8, 0, 0, 0, 0, time.UTC),
		Conditions: models.Conditions{
			Age: models.AgeRange{
				Start: 18,
				End:   40,
			},
			Gender:   "M",
			Country:  []string{"TW", "US"},
			Platform: []string{"android", "ios"},
		},
	}

	payload, err := json.Marshal(ad)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/ad", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	createAdHandler(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestListAdsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/ad", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	listAdsHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
