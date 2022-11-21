package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FormData struct {
	Name    string
	Age     string
	Country string
}

func main() {
	http.HandleFunc("/form", handleForm) 

	http.ListenAndServe(":3000", nil)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	var formData FormData
	_ = json.NewDecoder(r.Body).Decode(&formData)

	json.NewEncoder(w).Encode(formData)

	fmt.Fprintf(w, "hello")
}