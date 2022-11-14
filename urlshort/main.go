package main

import (
	"fmt"
	"net/http"
	"workspace/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	json := `[{"path": "/urlshort", "url": "https://github.com/gophercises/urlshort"}]`

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	handleErr(err)
	jsonHandler, err := handler.JSONHandler([]byte(json), mapHandler)
	handleErr(err)
	
	var answer string

	fmt.Scanf("%s\n", &answer)
	
	fmt.Println("Starting the server on: 8080")
	if answer == "json" {
		http.ListenAndServe(":8080", jsonHandler)		
	} else if answer == "yaml" {
		http.ListenAndServe(":8080", yamlHandler)		
	} else {
		fmt.Printf("Plese type 'json' or 'yaml'")
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}