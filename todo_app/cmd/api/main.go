package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"ruhultodoapp/cmd/internal/env"
)

type application struct {
	config        config
	templateCache map[string]*template.Template
	errorLog      *log.Logger
	infoLog       *log.Logger
}
type config struct {
	port string
}

func main() {
	port := env.GetString("PORT", "8080")
	templateCache, err := newTemplateCache()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		log.Fatalf("Failed to create template cache: %v", err)
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
	}

	mux.HandleFunc("/", app.home)

	fmt.Printf("Server started in port <<<<%s >>>>\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
