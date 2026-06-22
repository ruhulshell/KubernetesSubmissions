package main

import (
	"fmt"
	"net/http"
)

func (app *application) getTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "get todo running")
}
