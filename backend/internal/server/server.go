package server

import (
	"net/http"
	"time"
)

// New Return a new server
func New(h http.Handler, address string, writeTimeout, readTimeout, idleTimeout int) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      h,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}
}
