package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-playground/form/v4"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"tyres.kz/internal/models"
	"tyres.kz/internal/services"
)

type Handler struct {
	userService services.UserServiceInterface
	postService services.PostServiceInterface
	formDecoder *form.Decoder
	logger      *log.Logger
}

type TemplateData struct {
	success string
	error   string
}

func NewHandler(userService services.UserServiceInterface, postService services.PostServiceInterface, logger *log.Logger) *Handler {
	return &Handler{
		userService: userService,
		postService: postService,
		formDecoder: form.NewDecoder(),
		logger:      logger,
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

type RegisterForm struct {
	Email          string `form:"email"`
	Password       string `form:"password"`
	RepeatPassword string `form:"repeat_password"`
	Username       string `form:"username"`
	Phone          string `form:"phone"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var form RegisterForm
	err := r.ParseForm()
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.formDecoder.Decode(&form, r.PostForm)

	log.Println(form.Email, form.Password, form.RepeatPassword, form.Username, form.Phone)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Email:          form.Email,
		HashedPassword: string(hashedPassword),
		Username:       form.Username,
		Phone:          form.Phone,
	}

	_, err = h.userService.Create(user)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "login.tmpl.html"),
		filepath.Join("cmd/web/ui/views", "base-100vh.tmpl.html"),
	)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := TemplateData{
		success: "You have successfully registered!",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "profile.tmpl.html"),
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
	postId := httprouter.ParamsFromContext(r.Context()).ByName("id")

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
