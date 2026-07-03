package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	counter = 0
	mu      sync.Mutex
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", pingHandler)

	http.ListenAndServe(":8080", mux)

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	count := counter
	mu.Unlock()

	fmt.Fprintf(w, "pong %d", count)
}
