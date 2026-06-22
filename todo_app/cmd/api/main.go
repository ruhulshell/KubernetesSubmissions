package main

import (
	"fmt"
	"log"
	"net/http"
	"ruhultodoapp/cmd/internal/env"
)

type application struct {
	config config
}
type config struct {
	port string
}

func main() {
	port := env.GetString("PORT", "8080")
	mux := http.NewServeMux()
	cnf := config{
		port: port,
	}
	app := &application{
		config: cnf,
	}

	mux.HandleFunc("/", app.getTodo)

	fmt.Printf("Server started in port %s\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
