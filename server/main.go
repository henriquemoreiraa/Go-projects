package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Exemplo sem o gorilla/mux

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/haha", haha)

	http.ListenAndServe(":3000", nil)
}

// Exemplo com o gorilla/mux

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/haha", haha).Methods("GET")
	r.HandleFunc("/books/{title}/page/{page}", getBook).Methods("GET")
	
	http.ListenAndServe(":3000", r)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Your book is: %s and the page is: %s", vars["title"], vars["page"])
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello for all the 8 billion people on earth!")
}

func haha(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hahahaha so funny!")
}