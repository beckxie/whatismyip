package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetIP_PlainText(t *testing.T) {
	handler := NewHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/ip", nil)
	req.RemoteAddr = "192.168.1.100:12345"
	w := httptest.NewRecorder()

	handler.GetIP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("expected Content-Type text/plain, got %s", contentType)
	}

	body := w.Body.String()
	if body != "192.168.1.100" {
		t.Errorf("expected IP 192.168.1.100, got %s", body)
	}
}

func TestHandler_GetIP_JSON(t *testing.T) {
	handler := NewHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/ip?format=json", nil)
	req.RemoteAddr = "192.168.1.100:12345"
	req.Header.Set("User-Agent", "test-agent")
	w := httptest.NewRecorder()

	handler.GetIP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var resp IPResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp.IP != "192.168.1.100" {
		t.Errorf("expected IP 192.168.1.100, got %s", resp.IP)
	}

	if resp.Version != "IPv4" {
		t.Errorf("expected version IPv4, got %s", resp.Version)
	}

	if resp.Request.UserAgent != "test-agent" {
		t.Errorf("expected user agent test-agent, got %s", resp.Request.UserAgent)
	}
}

func TestHandler_GetIP_CORS(t *testing.T) {
	handler := NewHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/ip", nil)
	req.RemoteAddr = "192.168.1.100:12345"
	w := httptest.NewRecorder()

	handler.GetIP(w, req)

	corsHeader := w.Header().Get("Access-Control-Allow-Origin")
	if corsHeader != "*" {
		t.Errorf("expected CORS header *, got %s", corsHeader)
	}
}

func TestHandler_GetIP_Preflight(t *testing.T) {
	handler := NewHandler()

	req := httptest.NewRequest(http.MethodOptions, "/api/ip", nil)
	w := httptest.NewRecorder()

	handler.GetIP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204 for preflight, got %d", w.Code)
	}
}
