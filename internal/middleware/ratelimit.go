// Package middleware provides HTTP middleware components.
package middleware

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/beckxie/whatismyip/internal/ip"
)

// ipLimiter wraps a rate.Limiter with last access time for TTL cleanup.
type ipLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimiter is a per-IP rate limiting middleware with automatic cleanup.
type RateLimiter struct {
	limiters sync.Map
	rate     rate.Limit
	burst    int
	ttl      time.Duration
}

// NewRateLimiter creates a new RateLimiter.
// r is the rate of requests per second, b is the burst size.
// Starts a background goroutine to clean up expired entries.
func NewRateLimiter(r float64, b int) *RateLimiter {
	rl := &RateLimiter{
		rate:  rate.Limit(r),
		burst: b,
		ttl:   30 * time.Minute, // Default TTL: 30 minutes
	}

	// Start background cleanup goroutine
	go rl.cleanupLoop(10 * time.Minute)

	return rl
}

// getLimiter returns the rate limiter for the given IP.
func (rl *RateLimiter) getLimiter(clientIP string) *rate.Limiter {
	now := time.Now()

	if v, exists := rl.limiters.Load(clientIP); exists {
		il := v.(*ipLimiter)
		il.lastSeen = now
		return il.limiter
	}

	il := &ipLimiter{
		limiter:  rate.NewLimiter(rl.rate, rl.burst),
		lastSeen: now,
	}
	rl.limiters.Store(clientIP, il)
	return il.limiter
}

// cleanupLoop periodically removes expired entries.
func (rl *RateLimiter) cleanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		rl.cleanup()
	}
}

// cleanup removes entries that haven't been accessed within TTL.
func (rl *RateLimiter) cleanup() {
	now := time.Now()
	rl.limiters.Range(func(key, value any) bool {
		il := value.(*ipLimiter)
		if now.Sub(il.lastSeen) > rl.ttl {
			rl.limiters.Delete(key)
		}
		return true
	})
}

// Limit is the middleware handler.
func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := ip.GetIP(r)
		limiter := rl.getLimiter(clientIP)

		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
