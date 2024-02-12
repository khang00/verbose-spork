package handler

import (
	"github.com/khang00/verbose-spork/internal/pkg/jwt"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

const JWTUserNameKey = "username"
const JWTUserIDKey = "user_id"

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if len(authHeader) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized token is empty"))
			return
		}

		token, err := jwt.ValidateToken(strings.Split(authHeader[0], " ")[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized error parsing the JWT"))
			return
		}

		claims, _ := jwt.GetClaims(token)

		ctx := injectClaimsToCtx(r.Context(), claims)
		endpointHandler(w, r.WithContext(ctx))
	}
}

func injectClaimsToCtx(ctx context.Context, claims map[string]interface{}) context.Context {
	for k, v := range claims {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}

func GenerateJWTToken(username string, userID string) (string, error) {
	return jwt.CreateToken(map[string]string{
		JWTUserNameKey: username,
		JWTUserIDKey:   userID,
	})
}
