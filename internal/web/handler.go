package web

import (
	"log/slog"
	"net/http"
	"text/template"

	"github.com/beckxie/whatismyip/internal/ip"
)

// TemplateData holds data for the HTML template.
type TemplateData struct {
	IPInfo        string
	RequestHeader *http.Header
}

// Handler handles web page requests.
type Handler struct {
	tpl *template.Template
}

// NewHandler creates a new web handler with the given template path.
func NewHandler(tmplPath string) (*Handler, error) {
	tpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}
	return &Handler{tpl: tpl}, nil
}

// Index handles the main page request.
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	clientIP := ip.GetIP(r)

	data := &TemplateData{
		IPInfo:        "Can't identify your IP address.",
		RequestHeader: &r.Header,
	}

	if clientIP != "" {
		data.IPInfo = "Your IP address: " + clientIP
	}

	if err := h.tpl.Execute(w, data); err != nil {
		slog.Error("failed to execute template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
