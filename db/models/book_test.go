package models

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)
func init() {
	Init()
}

func TestGetBooks(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooksController)
	req, _ := http.NewRequest("GET", "/books", nil)

	res := executeRequest(req, r)
	checkResponse(t, http.StatusOK, res.Code)

	if body := res.Body.String(); body == "[]" {
		t.Errorf("Expected some data. Got %s", body)
	}
}

func TestGetBook(t *testing.T) {
	if _, err := GetBook(1); err == nil {
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

func checkResponse(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
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

func executeRequest(req *http.Request, r *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

// func clearTable() {
// 	db.Exec("DELETE FROM books")
// }
