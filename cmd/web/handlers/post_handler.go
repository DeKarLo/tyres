package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"tyres.kz/internal/services"
)

type PostHandler struct {
	postService *services.PostService
	userService *services.UserService
	logger      *log.Logger
}

func (postHandler *PostHandler) CreatePostPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "create-post.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"))

	if err != nil {
		postHandler.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		postHandler.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (postHandler *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	return
}

func (postHandler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	return
}

func (postHandler *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	return
}

func (postHandler *PostHandler) UpdatePostPage(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("id")
	postIdInt, err := strconv.Atoi(postId)

	if err != nil {
		postHandler.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	post, err := postHandler.postService.GetPostByID(postIdInt)

	if err != nil {
		postHandler.logger.Println(err)
		http.Error(w, "Post was not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "update-post.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"post": post})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (postHandler *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	return
}

func (postHandler *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	return
}
