package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/favicon.ico")
}
func humansTxt(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/humans.txt")
}
func robotsTxt(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/robots.txt")
}

func main() {
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/humans.txt", humansTxt)
	http.HandleFunc("/robots.txt", robotsTxt)
	http.HandleFunc("/", greet)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
