package auth

import (
	"context"
	"fmt"
	"github.com/khang00/verbose-spork/internal/handler"
	"strconv"
)

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
