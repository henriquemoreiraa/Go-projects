package models

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)
func init() {
	Init()
}

func TestGetBooksEmpty(t *testing.T) {
	clearTable()

	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooksController)

	res := executeRequest(t, r, nil, "/books", "GET")
	checkResponse(t, http.StatusOK, res.Code)

	if body := res.Body.String(); len(body) != 3 {
		t.Errorf("Expected empty array. Got %s", body)
	}
}

func TestCreateBook(t *testing.T) {
	book := RequestBodyTest(t)

	r := mux.NewRouter()
	r.HandleFunc("/book/create", CreateBookController)
	
	res := executeRequest(t, r, bytes.NewReader(book), "/book/create", "POST")
	checkResponse(t, http.StatusOK, res.Code)
}

func TestGetBooks(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooksController)

	res := executeRequest(t, r, nil, "/books", "GET")
	checkResponse(t, http.StatusOK, res.Code)

	if body := res.Body.String(); len(body) == 3 {
		t.Errorf("Expected some data. Got %s", body)
	}
}

func TestGetBook(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/book/{id}", GetBooksController)

	res := executeRequest(t, r, nil, "/book/1", "GET")
	checkResponse(t, http.StatusOK, res.Code)

	if body := res.Body.String(); len(body) == 3 {
		t.Errorf("Expected some data. Got %s", body)
	}
}

func TestUpdateBook(t *testing.T) {
	book := RequestBodyTest(t)

	r := mux.NewRouter()
	r.HandleFunc("/book/update/{id}", UpdateBookController)
	
	res := executeRequest(t, r, bytes.NewReader(book), "/book/update/1", "PUT")
	checkResponse(t, http.StatusOK, res.Code)
}

func TestDeleteBook(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/book/delete/{id}", DeleteBookController)
	
	res := executeRequest(t, r, nil, "/book/delete/1", "DELETE")
	checkResponse(t, http.StatusOK, res.Code)
}

func checkResponse(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(t *testing.T, r *mux.Router, body io.Reader, route, method string) *httptest.ResponseRecorder {

	req, err := http.NewRequest(method, route, body)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	return res
}

func RequestBodyTest(t *testing.T) []byte {
	book := Book{Title: "test", Description: "test", Author: "test", Pages: 1}

	j, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	return j
}

func clearTable() {
	db.Exec("DELETE FROM books")
	db.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}
