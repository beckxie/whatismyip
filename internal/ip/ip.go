package ip

import (
	"net"
	"net/http"
	"strings"
)

// GetIPVersion returns the IP version string ("IPv4" or "IPv6").
func GetIPVersion(ip string) string {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return ""
	}
	if parsedIP.To4() != nil {
		return "IPv4"
	}
	return "IPv6"
}

// GetIPWithVersion returns the client IP and its version.
func GetIPWithVersion(r *http.Request) (ip string, version string) {
	ip = GetIP(r)
	version = GetIPVersion(ip)
	return
}

// GetIP returns the client's public IP address.
// It checks X-Forwarded-For, X-Real-IP, and falls back to RemoteAddr.
func GetIP(r *http.Request) string {
	// 1. Check X-Forwarded-For (standard for multiple proxies)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		if clientIP := strings.TrimSpace(ips[0]); clientIP != "" {
			if parsedIP := net.ParseIP(clientIP); parsedIP != nil {
				return parsedIP.String()
			}
		}
	}

	// 2. Check X-Real-IP (often set by simple proxies like Nginx/Caddy)
	if xip := r.Header.Get("X-Real-IP"); xip != "" {
		if parsedIP := net.ParseIP(strings.TrimSpace(xip)); parsedIP != nil {
			return parsedIP.String()
		}
	}

	// 3. Fallback to RemoteAddr (direct connection)
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		if parsedIP := net.ParseIP(host); parsedIP != nil {
			return parsedIP.String()
		}
	}

	return ""
}
