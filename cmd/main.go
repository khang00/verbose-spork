package main

import (
	"fmt"
	"github.com/khang00/verbose-spork/internal/handler"
	"github.com/khang00/verbose-spork/internal/handler/auth"
	"github.com/khang00/verbose-spork/internal/handler/keyword"
	"github.com/khang00/verbose-spork/internal/store"
	"net/http"
)

func main() {
	db, err := store.NewPostgresStore(nil)
	if err != nil {
		fmt.Println(err)
	}

	authService := auth.NewAuthHandler(db)
	keywordService := keyword.NewKeywordHandler()

	setupHandler(authService, keywordService)

	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setupHandler(auth *auth.AuthHandler, keyword *keyword.KeywordHandler) {
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/api/user/signup", auth.Signup)
	http.HandleFunc("/api/user/signin", auth.Signin)

	http.HandleFunc("/api/keyword/upload", handler.VerifyJWT(keyword.UploadKeywords))
}
