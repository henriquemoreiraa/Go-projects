package main

import (
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/henriquemoreiraa/Gophercises/models"
	"github.com/henriquemoreiraa/Gophercises/models"
	"github.com/henriquemoreiraa/Gophercises/routes"
)

func main() {
	r := mux.NewRouter()
	models.Init()
	routes.Routes(r)

	http.ListenAndServe(":8080", r)
}
