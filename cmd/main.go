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

	authHandler := auth.NewAuthHandler(db)
	keywordHandler := keyword.NewKeywordHandler(db)

	setupHandler(authHandler, keywordHandler)

	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setupHandler(authhandler *auth.AuthHandler, keywordHandler *keyword.KeywordHandler) {
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/api/user/signup", authhandler.Signup)
	http.HandleFunc("/api/user/signin", authhandler.Signin)

	http.HandleFunc("/api/keyword/upload", handler.VerifyJWT(keywordHandler.UploadKeywords))
}
