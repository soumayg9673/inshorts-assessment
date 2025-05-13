package v1hdr

import "net/http"

func (h *V1Hdr) news() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /category", h.GetNewsByCategory)

	h.MUX.Handle("/news/", http.StripPrefix("/news", mux))
}

func (h *V1Hdr) GetNewsByCategory(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
