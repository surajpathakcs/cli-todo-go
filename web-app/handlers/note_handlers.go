package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// Dummy data for notes (this could be replaced with database data)
var notes = []struct {
	ID    string
	Title string
}{
	{"1", "First Note"},
	{"2", "Second Note"},
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	// Render the template with the notes data
	tmpl.Execute(w, struct {
		Notes []struct {
			ID    string
			Title string
		}
	}{notes})
}

func NewNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render new note page (HTML form)
		tmpl, err := template.ParseFiles("templates/new_note.html")
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Handle form submission to create a new note
		// Here you could add the new note to the database or data structure
		fmt.Fprintln(w, "Note Created")
	}
}

func EditNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Dummy note data (replace with data from your DB or data structure)
	note := struct {
		ID      string
		Title   string
		Content string
	}{id, "Sample Note", "This is a sample note content"}

	if r.Method == "GET" {
		// Render the note editing page (HTML form)
		tmpl, err := template.ParseFiles("templates/edit_note.html")
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, note)
	} else if r.Method == "POST" {
		// Handle the POST request to update the note
		// Save the edited note (e.g., update database)
		fmt.Fprintf(w, "Note with ID %s Updated", id)
	}
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Dummy note data (replace with data from your DB or data structure)
	note := struct {
		ID    string
		Title string
	}{id, "Sample Note"}

	if r.Method == "GET" {
		// Render the delete confirmation page
		tmpl, err := template.ParseFiles("templates/delete_note.html")
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, note)
	} else if r.Method == "POST" {
		// Handle deleting the note
		// (Usually, we would delete from a database here)
		fmt.Fprintf(w, "Note with ID %s Deleted", id)
	}
}
