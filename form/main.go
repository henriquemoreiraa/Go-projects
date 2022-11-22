package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FormData struct {
	Name    string `json:"name"`
	Age     int `json:"age"`
	Country string `json:"country"`
}

func main() {
	http.HandleFunc("/decode", decode) 
	http.HandleFunc("/encode", encode) 

	http.ListenAndServe(":3000", nil)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var formData FormData
	_ = json.NewDecoder(r.Body).Decode(&formData)

	fmt.Fprintf(w, "Your name is %s you are %d years old and you live in %s", formData.Name, formData.Age, formData.Country)
}

func encode(w http.ResponseWriter, r *http.Request) {
	user1 := FormData{
		Name: "Jose",
		Age: 20,
		Country: "France",
	}

	json.NewEncoder(w).Encode(user1)
}