package v1hdr

import (
	"net/http"
	"strconv"

	"github.com/soumayg9673/inshorts-assessment/cmd/api/errors"
	cjson "github.com/soumayg9673/inshorts-assessment/cmd/api/json"
)

func (h *V1Hdr) news() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /category", h.GetNewsByCategory)
	mux.HandleFunc("GET /score", h.GetNewsByScore)
	mux.HandleFunc("GET /source", h.GetNewsBySource)

	h.MUX.Handle("/news/", http.StripPrefix("/news", mux))
}

func (h *V1Hdr) GetNewsByCategory(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	/*
		Validate key "cat" in query parameters
		Return statusCode = 400 BAD REQUEST
		1. if "cat" not found in query parameters
		2. Key "cat" have no value
	*/
	if v, ok := q["query"]; len(v) == 0 || !ok {
		cjson.WriteJSONError(w, errors.BadRequst())
		return
	}

	data, err := h.SVC.GetNewsByCategory(q["query"])
	if err != nil {
		cjson.WriteJSONError(w, errors.SomethingWentWrong())
		return
	}
	cjson.WriteJSON(w, 200, data)
}

func (h *V1Hdr) GetNewsByScore(w http.ResponseWriter, r *http.Request) {
	data, err := h.SVC.GetNewsByScore()
	if err != nil {
		cjson.WriteJSONError(w, errors.SomethingWentWrong())
		return
	}
	cjson.WriteJSON(w, 200, data)
}

func (h *V1Hdr) GetNewsBySource(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	/*
		Validate key "query" in query parameters
		Return statusCode = 400 BAD REQUEST
		1. if "query" not found in query parameters
		2. Key "query" have no value
	*/
	if v, ok := q["query"]; len(v) == 0 || !ok {
		cjson.WriteJSONError(w, errors.BadRequst())
		return
	}

	// Check query data type for source (int)
	source, err := strconv.Atoi(q["query"][0])
	if err != nil {
		cjson.WriteJSONError(w, errors.BadRequst())
		return
	}

	data, err := h.SVC.GetNewsBySource(source)
	if err != nil {
		cjson.WriteJSONError(w, errors.SomethingWentWrong())
		return
	}
	cjson.WriteJSON(w, 200, data)
}
