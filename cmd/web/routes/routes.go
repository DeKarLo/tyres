package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"tyres.kz/cmd/web/handlers"
)

func Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc("GET", "/", handlers.Home)
	router.HandlerFunc("GET", "/about", handlers.About)
	router.HandlerFunc("GET", "/signup", handlers.RegisterPage)
	router.HandlerFunc("POST", "/signup", handlers.Register)
	router.HandlerFunc("GET", "/login", handlers.LoginPage)
	router.HandlerFunc("POST", "/login", handlers.Login)
	router.HandlerFunc("GET", "/logout", handlers.Logout)
	router.HandlerFunc("GET", "/profile", handlers.User)
	router.HandlerFunc("GET", "/post/create", handlers.CreatePostPage)
	router.HandlerFunc("GET", "/post/view/:id", handlers.GetPost)
	router.HandlerFunc("POST", "/post/", handlers.CreatePost)
	router.HandlerFunc("GET", "/posts", handlers.GetAllPosts)
	router.HandlerFunc("PUT", "/post/update/:id", handlers.UpdatePost)
	router.HandlerFunc("DELETE", "/post/delete/:id", handlers.DeletePost)

	return router
}
