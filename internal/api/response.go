package api

import "net/http"

// IPResponse represents the JSON response for the /api/ip endpoint.
type IPResponse struct {
	IP        string      `json:"ip"`
	Version   string      `json:"version"`
	Network   NetworkInfo `json:"network"`
	Request   RequestInfo `json:"request"`
	Proxy     ProxyInfo   `json:"proxy"`
	Timestamp string      `json:"timestamp"`
}

// NetworkInfo contains IPv4 and IPv6 addresses.
type NetworkInfo struct {
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6,omitempty"`
}

// RequestInfo contains HTTP request metadata.
type RequestInfo struct {
	Method         string `json:"method"`
	Path           string `json:"path"`
	UserAgent      string `json:"user_agent"`
	Accept         string `json:"accept,omitempty"`
	AcceptLanguage string `json:"accept_language,omitempty"`
	AcceptEncoding string `json:"accept_encoding,omitempty"`
}

// ProxyInfo contains reverse proxy related headers.
type ProxyInfo struct {
	Detected       bool   `json:"detected"`
	XForwardedFor  string `json:"x_forwarded_for,omitempty"`
	XRealIP        string `json:"x_real_ip,omitempty"`
	CFConnectingIP string `json:"cf_connecting_ip,omitempty"`
}

// NewIPResponse creates an IPResponse from an HTTP request.
func NewIPResponse(r *http.Request, clientIP, ipVersion, timestamp string) *IPResponse {
	xff := r.Header.Get("X-Forwarded-For")
	xrip := r.Header.Get("X-Real-IP")
	cfip := r.Header.Get("CF-Connecting-IP")

	resp := &IPResponse{
		IP:        clientIP,
		Version:   ipVersion,
		Timestamp: timestamp,
		Network: NetworkInfo{
			IPv4: "",
			IPv6: "",
		},
		Request: RequestInfo{
			Method:         r.Method,
			Path:           r.URL.Path,
			UserAgent:      r.Header.Get("User-Agent"),
			Accept:         r.Header.Get("Accept"),
			AcceptLanguage: r.Header.Get("Accept-Language"),
			AcceptEncoding: r.Header.Get("Accept-Encoding"),
		},
		Proxy: ProxyInfo{
			Detected:       xff != "" || xrip != "" || cfip != "",
			XForwardedFor:  xff,
			XRealIP:        xrip,
			CFConnectingIP: cfip,
		},
	}

	// Set IPv4 or IPv6 based on version
	if ipVersion == "IPv4" {
		resp.Network.IPv4 = clientIP
	} else if ipVersion == "IPv6" {
		resp.Network.IPv6 = clientIP
	}

	return resp
}
