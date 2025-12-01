package transport

import (
	"encoding/json"
	"net/http"
)

type HTTPHandlers struct {
	//
}

func NewHTTPHandlers() *HTTPHandlers {
	return &HTTPHandlers{
		//
	}
}

func (h *HTTPHandlers) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
