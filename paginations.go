package main

import (
	"fmt"
	"net/http"
)

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
