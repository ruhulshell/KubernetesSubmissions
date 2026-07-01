package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	randomStrLog string
)

func main() {
	randomStr := generateUUID()
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go logOutput(randomStr, ctx, &wg)

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(randomStrLog))

	})

	errr := http.ListenAndServe(":3070", nil)
	cancel()
	wg.Wait()
	if errr != nil {
		log.Fatal(errr)
	}
}

func logOutput(str string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
			temp := fmt.Sprintf("%s: %s\n", timestamp, str)
			log.Println(temp)
			randomStrLog = temp
		}
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
