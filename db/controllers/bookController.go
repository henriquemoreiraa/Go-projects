package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/henriquemoreiraa/Gophercises/models"
	"github.com/henriquemoreiraa/Gophercises/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	writeResponse(w, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvStringToInt(r)
	if err != nil {
		log.Fatal(err)
	}

	book, err := models.GetBook(id)
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook, err := models.CreateBook(r)
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, createBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvStringToInt(r)
	if err != nil {
		log.Fatal(err)
	}

	book, err := models.UpdateBook(r, id)
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvStringToInt(r)
	if err != nil {
		log.Fatal(err)
	}

	book, err := models.DeleteBook(id)
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, book)
}

func writeResponse[B BookConst](w http.ResponseWriter, b B) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

type BookConst interface {
	models.Book | []models.Book
}
