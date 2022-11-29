package routes

import (
	"github.com/gorilla/mux"
	"github.com/henriquemoreiraa/Gophercises/models"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/books", models.GetBooksController).Methods("GET")
	r.HandleFunc("/book/{id}", models.GetBookController).Methods("GET")
	r.HandleFunc("/book/create", models.CreateBookController).Methods("POST")
	r.HandleFunc("/book/update/{id}", models.UpdateBookController).Methods("PUT")
	r.HandleFunc("/book/delete/{id}", models.DeleteBookController).Methods("DELETE")
}