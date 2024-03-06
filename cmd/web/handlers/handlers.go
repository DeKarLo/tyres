package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"time"

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
	Success string
	Error   string
	Posts   []*models.Post
	User    *models.User
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

	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Email:          form.Email,
		HashedPassword: string(form.Password),
		Username:       form.Username,
		Phone:          form.Phone,
	}

	_, err = h.userService.Create(user)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "login.tmpl.html"),
	))

	// if err != nil {
	// 	h.logger.Println(err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	data := TemplateData{
		Success: "You have successfully registered!",
		Error:   "",
	}

	err = tmpl.Execute(w, &data)

	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "login.tmpl.html"),
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

type loginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var form loginForm
	err := r.ParseForm()
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.formDecoder.Decode(&form, r.PostForm)

	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user, err := h.userService.GetByEmail(form.Email)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	h.logger.Println(user.HashedPassword)
	h.logger.Println(form.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(form.Password))
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:    "UserID",
		Value:   strconv.Itoa(user.ID),
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "UserID",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}

// User profile handler

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "profile.tmpl.html"),
	))

	cookie, _ := r.Cookie("UserID")

	id, _ := strconv.Atoi(cookie.Value)

	user, _ := h.userService.GetByID(id)

	// if err != nil {
	// 	h.logger.Println(err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	h.logger.Println(user)
	err := tmpl.Execute(w, &user)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Post handlers

type PostForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	Img     string `form:"img"`
	Price   int    `form:"price"`
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var form PostForm
	cookie, _ := r.Cookie("UserID")
	userID, _ := strconv.Atoi(cookie.Value)
	user, _ := h.userService.GetByID(userID)
	err := r.ParseForm()
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	post := &models.Post{
		Title:   form.Title,
		Content: form.Content,
		Img:     form.Img,
		Price:   form.Price,
		UserID:  user.ID,
	}

	err = h.postService.CreatePost(post)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/posts", http.StatusFound)
}

func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var id int
	var user *models.User
	cookie, err := r.Cookie("UserID")
	if err != nil {
		h.logger.Println("No UserID cookie found:", err)
	} else {
		id, err = strconv.Atoi(cookie.Value)
		if err != nil {
			h.logger.Println(err)
		} else {
			user, err = h.userService.GetByID(id)
			if err != nil {
				h.logger.Println(err)
				user = nil
			}
		}
	}

	templateData := TemplateData{
		Posts: posts,
		User:  user,
	}

	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "posts.tmpl.html"))
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, templateData)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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
		filepath.Join("cmd/web/ui/views/pages", "update-post.tmpl.html"))

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
	postId := httprouter.ParamsFromContext(r.Context()).ByName("id")

	postIdInt, err := strconv.Atoi(postId)
	h.logger.Println(postIdInt)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.postService.DeletePost(postIdInt)

	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/posts", http.StatusFound)
}

// Home and about page handlers

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	var id int
	var user *models.User
	cookie, err := r.Cookie("UserID")
	if err != nil {
		h.logger.Println("No UserID cookie found:", err)
	} else {
		id, err = strconv.Atoi(cookie.Value)
		if err != nil {
			h.logger.Println(err)
		} else {
			user, err = h.userService.GetByID(id)
			if err != nil {
				h.logger.Println(err)
				user = nil
			}
		}
	}

	templateData := TemplateData{
		User: user,
	}

	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "home.tmpl.html"),
	)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, templateData)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {

	var id int
	var user *models.User
	cookie, err := r.Cookie("UserID")
	if err != nil {
		h.logger.Println("No UserID cookie found:", err)
	} else {
		id, err = strconv.Atoi(cookie.Value)
		if err != nil {
			h.logger.Println(err)
		} else {
			user, err = h.userService.GetByID(id)
			if err != nil {
				h.logger.Println(err)
				user = nil
			}
		}
	}

	templateData := TemplateData{
		User: user,
	}
	tmpl, err := template.ParseFiles(
		filepath.Join("cmd/web/ui/views/pages", "about.tmpl.html"),
	)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, templateData)
	if err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
