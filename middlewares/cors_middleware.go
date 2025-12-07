package middlewares

import (
	"net/http"
)

type CorsMiddleware struct {
	handler http.Handler
}

func NewCorsMiddleware(handler http.Handler) *CorsMiddleware {
	return &CorsMiddleware{handler}
}

func (m *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	// Allowed origins for both local dev and production
	allowedOrigins := map[string]bool{
		"http://localhost:3000":               true,
		"https://your-frontend.vercel.app":    true,
	}

	// If the origin is in our allowlist, allow it
	if allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Vary", "Origin") // important!
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	m.handler.ServeHTTP(w, r)
}
