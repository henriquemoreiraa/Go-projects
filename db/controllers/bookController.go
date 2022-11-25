package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/henriquemoreiraa/Gophercises/models"
	"github.com/henriquemoreiraa/Gophercises/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	writeResponse(w, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := utils.ConvStringToInt(r)

	book := models.GetBook(id)	
	writeResponse(w, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := models.CreateBook(r)
	writeResponse(w, createBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := utils.ConvStringToInt(r)

	book := models.UpdateBook(r, id)
	writeResponse(w, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := utils.ConvStringToInt(r)

	book := models.DeleteBook(id)
	writeResponse(w, book)
}

func writeResponse[B BookConst](w http.ResponseWriter, b B ) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

type BookConst interface {
	models.Book | []models.Book
}
