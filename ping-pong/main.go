package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	counter = 0
	mu      sync.Mutex
)

func main() {
	logDir := os.Getenv("LOG_DIRECTORY")
	if logDir == "" {
		logDir = "files"
	}

	filePath := filepath.Join(logDir, "counter.log")

	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		fmt.Printf("Writer: Failed to create directory: %v\n", err)
		return
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		initialData := []byte("0\n")
		err = os.WriteFile(filePath, initialData, 0644)
		if err != nil {
			fmt.Printf("Writer: Failed to initialize counter.log on startup: %v\n", err)
			return
		}
		fmt.Println("Writer: counter.log not found. Created and initialized with 0.")
	} else {
		data, err := os.ReadFile(filePath)
		if err == nil {
			trimmed := strings.TrimSpace(string(data))
			if existingCount, parseErr := strconv.Atoi(trimmed); parseErr == nil {
				counter = existingCount
				fmt.Printf("Writer: Restored counter state to %d from storage\n", counter)
			}
		}
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		counter++

		logLine := fmt.Sprintf("%d\n", counter)

		err := os.WriteFile(filePath, []byte(logLine), 0644)
		if err != nil {
			fmt.Printf("Writer: Failed to write file: %v\n", err)
			http.Error(w, "Failed to write counter", http.StatusInternalServerError)
			counter--
			return
		}

		fmt.Fprintf(w, "pong %d", counter)
	})

	fmt.Println("Server starting on port :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
