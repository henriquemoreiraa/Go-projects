package routes

import (
	"github.com/gorilla/mux"
	"github.com/henriquemoreiraa/Gophercises/controllers"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/update/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/delete/{id}", controllers.DeleteBook).Methods("DELETE")
}