package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

)

// =========================== HELPER ===========================
func resolveID(r *http.Request, w http.ResponseWriter) int64 {
	rawID := chi.URLParam(r, "id")
	if rawID == "" { return 0 }

	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
	}

	return id
}

func resolvePage(r *http.Request) (int64, int64) {
	const (
		defaultPage = 1
		defaultSize = 5
		maxSize     = 50
	)

	rawPage := r.URL.Query().Get("page")
	rawSize := r.URL.Query().Get("size")

	page, err := strconv.ParseInt(rawPage, 10, 64)
	if err != nil || page <= 0 { page = defaultPage }

	size, err := strconv.ParseInt(rawSize, 10, 64)
	if err != nil || size <= 0 { size = defaultSize }

	if size > maxSize { size = maxSize }

	return page, size
}
// =========================== HELPER ===========================
