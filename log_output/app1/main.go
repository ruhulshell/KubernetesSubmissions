package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	logDir := os.Getenv("LOG_DIRECTORY")
	if logDir == "" {
		logDir = "files"
	}

	filePath := filepath.Join(logDir, "app.log")

	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		fmt.Printf("Writer: Failed to create directory: %v\n", err)
		return
	}

	randomStr := generateUUID()
	fmt.Printf("Writer started. Generating logs with UUID: %s\n", randomStr)

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Writer: Failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	for {
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		logLine := fmt.Sprintf("%s: %s\n", timestamp, randomStr)

		_, err := f.WriteString(logLine)
		if err != nil {
			fmt.Printf("Writer: Failed to write to file: %v\n", err)
		}

		f.Sync()
		time.Sleep(5 * time.Second)
	}
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "00000000-0000-0000-0000-000000000000"
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
