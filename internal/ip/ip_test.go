package ip

import (
	"net/http"
	"testing"
)

func TestGetIP(t *testing.T) {
	tests := []struct {
		name       string
		headers    map[string]string
		remoteAddr string
		expected   string
	}{
		{
			name:     "X-Forwarded-For: Single IP",
			headers:  map[string]string{"X-Forwarded-For": "1.2.3.4"},
			expected: "1.2.3.4",
		},
		{
			name:     "X-Forwarded-For: Multiple IPs",
			headers:  map[string]string{"X-Forwarded-For": "1.2.3.4, 5.6.7.8"},
			expected: "1.2.3.4",
		},
		{
			name:     "X-Real-IP fallback",
			headers:  map[string]string{"X-Real-IP": "5.6.7.8"},
			expected: "5.6.7.8",
		},
		{
			name:       "RemoteAddr fallback",
			remoteAddr: "9.10.11.12:12345",
			expected:   "9.10.11.12",
		},
		{
			name:     "Invalid XFF falls back to X-Real-IP",
			headers:  map[string]string{"X-Forwarded-For": "invalid", "X-Real-IP": "5.6.7.8"},
			expected: "5.6.7.8",
		},
		{
			name: "Preference order (XFF > X-Real-IP)",
			headers: map[string]string{
				"X-Forwarded-For": "1.2.3.4",
				"X-Real-IP":       "5.6.7.8",
			},
			expected: "1.2.3.4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			for k, v := range tt.headers {
				req.Header.Set(k, v)
			}
			if tt.remoteAddr != "" {
				req.RemoteAddr = tt.remoteAddr
			} else {
				req.RemoteAddr = "127.0.0.1:0" // Default
			}

			got := GetIP(req)
			if got != tt.expected {
				t.Errorf("GetIP() = %v, want %v", got, tt.expected)
			}
		})
	}
}
