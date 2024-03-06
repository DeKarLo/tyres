package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "signup.tmpl.html"),
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

func Register(w http.ResponseWriter, r *http.Request) {
	return
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "login.tmpl.html"),
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

func Login(w http.ResponseWriter, r *http.Request) {
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	return
}

func User(w http.ResponseWriter, r *http.Request) {
	return
}

