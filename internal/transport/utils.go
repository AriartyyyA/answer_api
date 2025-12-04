package transport

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSONError writes a standardized JSON error response.
func WriteJSONError(w http.ResponseWriter, message string, code int) {
	log.Printf("Error: %s (Code: %d)", message, code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
