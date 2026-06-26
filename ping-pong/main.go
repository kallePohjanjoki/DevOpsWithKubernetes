package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

var counter atomic.Int64

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		count := counter.Add(1)
		fmt.Fprintf(w, "pong %d", count-1)
	})

	fmt.Printf("Server started in port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}