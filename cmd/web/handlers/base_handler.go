package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "home.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "about.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
