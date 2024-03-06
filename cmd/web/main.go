package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"tyres.kz/cmd/web/handlers"
	"tyres.kz/cmd/web/routes"
	"tyres.kz/internal/repositories"
	"tyres.kz/internal/services"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := sql.Open("mysql", "web:123455@tcp(localhost:3306)/tyres")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	userRepo, _ := repositories.NewUserRepository(db, errorLog)
	postRepo, _ := repositories.NewPostRepository(db, errorLog)

	userService, _ := services.NewUserService(userRepo, errorLog)
	postService, _ := services.NewPostService(postRepo, errorLog)

	handlers := handlers.NewHandler(&userService, &postService, errorLog)

	srv := &http.Server{
		Addr:         ":4000",
		ErrorLog:     errorLog,
		Handler:      routes.Routes(handlers),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running on http://localhost:8080")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
