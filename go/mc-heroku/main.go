package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {

	go Sync()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)

	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}

	http.ListenAndServe(":"+port, mux)

}

// Sync to github
func Sync() {

	for {
		time.Sleep(time.Second * 600)
		cmd := exec.Command("/bin/sh", "-c", "/minecraft/sync.sh")
		if _, err := cmd.Output(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}

// Root send message to root
func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		t, err := template.ParseFiles("main.html")
		if err != nil {
			log.Printf("root failed: %e", err)
		}
		t.Execute(w, nil)

	} else {
		w.Write([]byte("Error"))
	}
}
