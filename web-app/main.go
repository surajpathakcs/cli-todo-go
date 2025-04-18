package main

import (
	"net/http"
	"web-app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files from /static/ folder
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("D:/code/go/web-app/static"))))

	// Routes
	r.HandleFunc("/", handlers.HandleHome).Methods("GET")
	r.HandleFunc("/note/new", handlers.NewNoteHandler).Methods("GET", "POST")
	r.HandleFunc("/note/{id}", handlers.EditNoteHandler).Methods("GET", "POST")
	r.HandleFunc("/note/{id}/delete", handlers.DeleteNoteHandler).Methods("POST")

	http.ListenAndServe(":8000", r)
}
