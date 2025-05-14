package v1hdr

import (
	"net/http"

	"github.com/soumayg9673/inshorts-assessment/cmd/api/errors"
	cjson "github.com/soumayg9673/inshorts-assessment/cmd/api/json"
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
		Return statusCode = 400 BAD REQUEST
		1. if "cat" not found in query parameters
		2. Key "cat" have no value
	*/
	if v, ok := q["cat"]; len(v) == 0 || !ok {
		cjson.WriteJSONError(w, errors.BadRequst())
		return
	}

	data, err := h.SVC.GetNewsByCategory(q["cat"])
	if err != nil {
		cjson.WriteJSONError(w, errors.SomethingWentWrong())
		return
	}
	cjson.WriteJSON(w, 200, data)
}
