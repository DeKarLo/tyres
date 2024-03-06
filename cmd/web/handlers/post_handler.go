package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func CreatePostPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "create-post.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"))

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

func GetPost(w http.ResponseWriter, r *http.Request) {
	return
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	return
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdatePostPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "update-post.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"))

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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	return
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	return
}
