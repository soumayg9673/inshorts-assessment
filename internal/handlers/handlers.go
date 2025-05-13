package handlers

import (
	"net/http"

	v1hdr "github.com/soumayg9673/inshorts-assessment/internal/handlers/v1"
	"github.com/soumayg9673/inshorts-assessment/internal/service"
)

func RegisterRoutes(mux *http.ServeMux, s service.Service) {
	v1hdr.NewV1Handlers(mux, s.V1)
}
