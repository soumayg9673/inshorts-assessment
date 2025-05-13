package v1hdr

import (
	"net/http"
)

func (h *V1Hdr) news() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /category", h.GetNewsByCategory)

	h.MUX.Handle("/news/", http.StripPrefix("/news", mux))
}

func (h *V1Hdr) GetNewsByCategory(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	/*
		Validate key "cat" in query parameters
		Return statusCode = 400 BAD REQUEST (if "cat" not found in query parameters)
	*/
	if _, ok := q["cat"]; !ok {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
}
