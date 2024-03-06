package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Message: "Hello, World!",
	}

	tmpl := template.Must(template.ParseFiles("test/template.tmpl.html"))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
