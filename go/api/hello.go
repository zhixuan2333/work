package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
)

// Mainpage write root page
func Mainpage(w http.ResponseWriter) {
	t, err := template.ParseFiles("main.html")
	if err != nil {
		log.Printf("root failed: %e", err)
	}
	t.Execute(w, nil)
}

// RemoteIP get ip
func RemoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("XRealIP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("XForwardedFor"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
