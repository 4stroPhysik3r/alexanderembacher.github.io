package main

import (
	"html/template"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePageHandler)

	log.Printf("Starting server at: http://localhost" + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	err := renderTempl(w, r)

	if err != nil {
		http.Error(w, err.Error(), 404)
	}
}

func renderTempl(w http.ResponseWriter, r *http.Request) error {
	templ, err := template.ParseFiles("index.html")

	if err != nil {
		return err
	}

	err = templ.Execute(w, nil)
	if err != nil {
		return err
	}

	return nil
}
