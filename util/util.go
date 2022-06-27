package util

import (
	"net/http"
)

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if len(ip) > 0 {
		return ip
	}
	ip = r.Header.Get("X-Real-Ip")
	if len(ip) > 0 {
		return ip
	}
	return r.RemoteAddr
}
