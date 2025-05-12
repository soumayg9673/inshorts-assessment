package v1hdr

import "net/http"

func (h V1Hdr) news() {
	mux := http.NewServeMux()
	// TODO: Register API routes
	h.MUX.Handle("/news/", http.StripPrefix("/news", mux))
}
