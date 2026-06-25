package main

import (
	"net/http"
	"time"
)

type Todo struct {
	ID     int
	Title  string
	Status string // "pending" or "completed"
}

type PageData struct {
	Todos          []Todo
	CompletedCount int
	Flash          string
	Year           int
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{ID: 1, Title: "Buy groceries", Status: "completed"},
		{ID: 2, Title: "Read 20 pages of a book", Status: "pending"},
		{ID: 3, Title: "Fix login page bug", Status: "completed"},
		{ID: 4, Title: "Reply to team emails", Status: "pending"},
		{ID: 5, Title: "Schedule dentist appointment", Status: "pending"},
	}

	data := PageData{
		Todos:          todos,
		CompletedCount: completedCount(todos),
		Flash:          "Task added successfully!", // set "" to hide
		Year:           time.Now().Year(),
	}

	app.render(w, "todos.tmpl", &templateData{data})

}
