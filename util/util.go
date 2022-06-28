package util

import (
	"crypto/md5"
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

func MD5Byte(data []byte) []byte {
	hash := md5.New()
	hash.Write(data)
	return hash.Sum(nil)
}
