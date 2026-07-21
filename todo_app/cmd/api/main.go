package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"ruhultodoapp/cmd/internal/downloader"
	"ruhultodoapp/cmd/internal/env"
)

type application struct {
	config        config
	templateCache map[string]*template.Template
	errorLog      *log.Logger
	infoLog       *log.Logger
	ImgCache      *downloader.ImgCache
}
type config struct {
	port string
}

func main() {
	logDir := os.Getenv("IMG_DIRECTORY")
	if logDir == "" {
		logDir = "images"
	}
	port := env.GetString("PORT", "8080")
	templateCache, err := newTemplateCache()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		log.Fatalf("Failed to create template cache: %v", err)
	}

	// image cache
	imgCache := &downloader.ImgCache{
		CacheDir:  logDir,
		ImagePath: logDir + "/image.jpg",
	}
	mux := http.NewServeMux()
	cnf := config{
		port: port,
	}
	app := &application{
		config:        cnf,
		templateCache: templateCache,
		errorLog:      errorLog,
		infoLog:       infoLog,
		ImgCache:      imgCache,
	}

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/image", app.ImgCache.ImageHandler)

	fmt.Printf("Server started in port <<<<%s >>>>\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
