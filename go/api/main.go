package main

import (
	"net/http"
	"os"
)

// Root Send msg to root
func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Mainpage(w)
	}
}

// Getip get ip page
func Getip(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		out := RemoteIP(r)
		w.Write([]byte(out))

	} else {
		w.Write([]byte("Err: the success method is GET"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)
	mux.HandleFunc("/getip", Getip)

	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}

	http.ListenAndServe(":"+port, mux)
}
