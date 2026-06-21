package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

func main() {
	randomStr := generateUUID()

	logOutput(randomStr)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		logOutput(randomStr)
	}
}

func logOutput(str string) {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	log.Printf("%s: %s\n", timestamp, str)
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "00000000-0000-0000-0000-000000000000"
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
