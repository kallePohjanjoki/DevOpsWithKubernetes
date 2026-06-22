package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	fmt.Printf("Server started in port NNNN")
	http.ListenAndServe(":"+port, nil)
}
