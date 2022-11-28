package models

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetBooks(t *testing.T) {
	if res := GetBooks(); res == nil {
		t.Error("test failed")
	}
}

func TestGetBook(t *testing.T) {
	if _, err := GetBook(1); err != nil {
		t.Error("test failed")
	}
}
func TestCreateBook(t *testing.T) {
	book, err := RequestBodyTest()
	if err != nil {
		t.Error(err)
	}

	r, _ := http.NewRequest("POST", "/book/create", bytes.NewReader(book))

	if _, err := CreateBook(r); err != nil {
		t.Error("test failed")
	}
}

func TestUpdateBook(t *testing.T) {
	book, err := RequestBodyTest()
	if err != nil {
		t.Error(err)
	}

	r, _ := http.NewRequest("POST", "/book/create", bytes.NewReader(book))

	if _, err := UpdateBook(r, 2); err != nil {
		t.Error("test failed")
	}
}

func TestDeleteBook(t *testing.T) {
	if _, err := DeleteBook(2); err != nil {
		t.Error("test failed")
	}
}

func RequestBodyTest() ([]byte, error) {
	book := Book{Title: "test", Description: "test", Author: "test", Pages: 1}

	j, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	return j, nil
}
