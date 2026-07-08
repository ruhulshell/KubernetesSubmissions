package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	logDir := os.Getenv("LOG_DIRECTORY")
	if logDir == "" {
		logDir = "files"
	}
	filePath := filepath.Join(logDir, "app.log")

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)

	})

	errr := http.ListenAndServe(":3070", nil)
	if errr != nil {
		log.Fatal(errr)
	}
}
