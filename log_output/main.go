package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"
)

var randomString string

func randomUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
	randomString = randomUUID()

	port := "8080"

	go func() {
		for {
			fmt.Printf("%s: %s\n", time.Now().UTC().Format(time.RFC3339Nano), randomString)
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s: %s", time.Now().UTC().Format(time.RFC3339Nano), randomString)
	})

	fmt.Printf("Server started in port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
