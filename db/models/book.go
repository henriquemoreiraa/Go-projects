package models

import (
	"encoding/json"
	"net/http"

	"github.com/henriquemoreiraa/Gophercises/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Pages       int    `json:"pages"`
}

var db *gorm.DB

func Init() {
	db = config.ConnectDb()
	db.AutoMigrate(&Book{})
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)

	return books
}

func GetBook(Id int64) (Book, error) {
	var book Book
	err := db.Where("ID=?", Id).Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func CreateBook(r *http.Request) (Book, error) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	err := db.Create(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func UpdateBook(r *http.Request, Id int64) (Book, error) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	json.NewDecoder(r.Body).Decode(&book)

	err := db.Save(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func DeleteBook(Id int64) (Book, error) {
	var book Book
	err := db.Where("ID=?", Id).Delete(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}
