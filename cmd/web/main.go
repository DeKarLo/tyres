package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"tyres.kz/cmd/web/routes"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	srv := &http.Server{
		Addr:         ":4000",
		ErrorLog:     errorLog,
		Handler:      routes.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running on http://localhost:8080")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
