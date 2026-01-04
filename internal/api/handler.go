package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/beckxie/whatismyip/internal/ip"
)

// Handler handles API requests.
type Handler struct{}

// NewHandler creates a new API handler.
func NewHandler() *Handler {
	return &Handler{}
}

// GetIP handles the /api/ip endpoint.
// Returns plain text IP by default, or JSON with ?format=json
func (h *Handler) GetIP(w http.ResponseWriter, r *http.Request) {
	clientIP, ipVersion := ip.GetIPWithVersion(r)

	// CORS headers for cross-origin requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Check format parameter
	format := r.URL.Query().Get("format")

	if format == "json" {
		w.Header().Set("Content-Type", "application/json")

		timestamp := time.Now().Format(time.RFC3339)
		resp := NewIPResponse(r, clientIP, ipVersion, timestamp)

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			slog.Error("failed to encode JSON response", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Default: return plain text IP
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, clientIP)
}
