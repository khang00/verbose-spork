package main

import (
	"fmt"
	"github.com/khang00/verbose-spork/internal/handler"
	"github.com/khang00/verbose-spork/internal/store"
	"net/http"
)

func main() {
	db, err := store.NewPostgresStore(nil)
	if err != nil {
		fmt.Println(err)
	}

	authService := handler.NewAuthHandler(db)

	setupHandler(authService)

	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setupHandler(auth *handler.AuthHandler) {
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/api/user/signup", auth.Signup)
	http.HandleFunc("/api/user/signin", auth.Signin)
}
