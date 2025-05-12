package v1hdr

import (
	"net/http"

	v1svc "github.com/soumayg9673/inshorts-assessment/internal/service/v1"
)

type V1Hdr struct {
	MUX *http.ServeMux
	SVC v1svc.V1
}

func NewV1Handlers(mux *http.ServeMux, svc v1svc.V1) {
	handlers := &V1Hdr{
		MUX: http.NewServeMux(),
		SVC: svc,
	}
	// Register API routes
	handlers.news() // route: /news
	mux.Handle("/v1/", http.StripPrefix("/v1", handlers.MUX))
}
