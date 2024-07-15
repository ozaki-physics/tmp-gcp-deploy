package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
