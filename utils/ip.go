package utils

import (
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {
	// 1. Check X-Forwarded-For (may contain multiple IPs)
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	// 2. Check X-Real-IP (common with Nginx)
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// 3. Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
