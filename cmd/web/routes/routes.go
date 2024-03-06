package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"tyres.kz/cmd/web/handlers"
)

func Routes(h *handlers.Handler) http.Handler {
	router := httprouter.New()

	router.HandlerFunc("GET", "/", h.Home)
	router.HandlerFunc("GET", "/about", h.About)
	router.HandlerFunc("GET", "/signup", h.RegisterPage)
	router.HandlerFunc("POST", "/signup", h.Register)
	router.HandlerFunc("GET", "/login", h.LoginPage)
	router.HandlerFunc("POST", "/login", h.Login)
	router.HandlerFunc("GET", "/logout", h.Logout)
	router.HandlerFunc("GET", "/profile", h.User)
	router.HandlerFunc("POST", "/post", h.CreatePost)
	router.HandlerFunc("GET", "/posts", h.GetAllPosts)
	router.HandlerFunc("PUT", "/post/update/:id", h.UpdatePost)
	router.HandlerFunc("GET", "/post/update/:id", h.UpdatePostPage)
	router.HandlerFunc("DELETE", "/post/delete/:id", h.DeletePost)

	return router
}
