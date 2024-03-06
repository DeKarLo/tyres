package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"tyres.kz/internal/services"
)

type Handler struct {
	userService *services.UserServiceInterface
	postService *services.PostServiceInterface
	logger 		*log.Logger
}

func NewHandler(userService *services.UserServiceInterface, postService *services.PostServiceInterface, logger *log.Logger) *Handler {
	return &Handler{
		userService: userService,
		postService: postService,
		logger: logger,
	}
}

// Authorization handlers

func (h *Handler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "signup.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"),
	)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "login.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"),
	)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	return
}

// User profile handler

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	return
}

// Post handlers

func (h *Handler) CreatePostPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "create-post.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"))

	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) UpdatePostPage(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("id")
	postIdInt, err := strconv.Atoi(postId)

	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	post, err := h.postService.GetPostByID(postIdInt)

	if err != nil {
		h.logger.Println(err)
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

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	return
}

// Home and about page handlers

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "home.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"),
	)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "about.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"),
	)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
