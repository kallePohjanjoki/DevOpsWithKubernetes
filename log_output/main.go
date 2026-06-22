package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	id := randomUUID()

	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().Format(time.RFC3339Nano), id)
		time.Sleep(5 * time.Second)
	}
}

func randomUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
