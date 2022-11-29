package models

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/henriquemoreiraa/Gophercises/utils"
)

func GetBooksController(w http.ResponseWriter, r *http.Request) {
	books := GetBooks()
	writeResponse(w, books)
}

func GetBookController(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvStringToInt(r)
	if err != nil {
		log.Fatal(err)
	}

	book, err := GetBook(id)
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, book)
}

func CreateBookController(w http.ResponseWriter, r *http.Request) {
	createBook, err := CreateBook(r)
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, createBook)
}

func UpdateBookController(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvStringToInt(r)
	if err != nil {
		log.Fatal(err)
	}

	book, err := UpdateBook(r, id)
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, book)
}

func DeleteBookController(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvStringToInt(r)
	if err != nil {
		log.Fatal(err)
	}

	book, err := DeleteBook(id)
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
	Book | []Book
}
