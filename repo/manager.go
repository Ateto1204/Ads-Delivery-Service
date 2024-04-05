package repo

import (
	"sync"

	"github.com/Ateto1204/Ads-Delivery-Service/models"
)

// Ads is an in-memory store for advertisements.
var Ads []models.Ad
var mu sync.Mutex

func addAd(ad models.Ad) {
	Ads = append(Ads, ad)
}

func muxLock() {
	mu.Lock()
}

func muxUnlock() {
	defer mu.Unlock()
}
