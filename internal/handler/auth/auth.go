package auth

import (
	"context"
	"fmt"
	"github.com/khang00/verbose-spork/internal/handler"
	"github.com/khang00/verbose-spork/internal/model"
	"strconv"
)

type UserStore interface {
	CreateUser(username string, password string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
}

func GetUserNameFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(handler.JWTUserIDKey)
	if val == nil {
		return "", fmt.Errorf("error: %s in context is empty %v", handler.JWTUserNameKey, val)
	}

	username, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("error: can not parse userID %v to string", val)
	}

	return username, nil
}

func GetUserIDFromContext(ctx context.Context) (uint, error) {
	val := ctx.Value(handler.JWTUserIDKey)
	if val == nil {
		return 0, fmt.Errorf("error: %s in context is empty %v", handler.JWTUserIDKey, val)
	}

	userIDStr, ok := val.(string)
	if !ok {
		return 0, fmt.Errorf("error: can not parse %v to string", val)
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, err
	}

	return uint(userID), nil
}

type AuthHandler struct {
	userStore UserStore
}

func NewAuthHandler(userStore UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}
