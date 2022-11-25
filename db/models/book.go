package models

import (
	"encoding/json"
	"net/http"

	"github.com/henriquemoreiraa/Gophercises/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string	`json:"title"`
	Description string	`json:"description"`
	Author string	`json:"author"`
	Pages int	`json:"pages"`
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

func GetBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)

	return book
}

func CreateBook(r *http.Request) Book {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	db.Create(&book)

	return book
}

func UpdateBook(r *http.Request, Id int64) Book {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	json.NewDecoder(r.Body).Decode(&book)	
	db.Save(&book)

	return book
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)

	return book
}