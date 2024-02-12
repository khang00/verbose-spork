package auth

import (
	"github.com/khang00/verbose-spork/internal/handler"
)

type AuthHandler struct {
	userStore handler.UserStore
}

func NewAuthHandler(userStore handler.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}
