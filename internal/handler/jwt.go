package handler

import (
	"github.com/khang00/verbose-spork/internal/pkg/jwt"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

const JWTUserNameKey = "username"

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header["Authorization"]
		if len(authHeader) == 0 {
			writer.WriteHeader(http.StatusUnauthorized)
			_, _ = writer.Write([]byte("Unauthorized token is empty"))
		}

		token, err := jwt.ValidateToken(strings.Split(authHeader[0], " ")[len(authHeader)-1])
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			_, _ = writer.Write([]byte("Unauthorized error parsing the JWT"))
		}

		claims, _ := jwt.GetClaims(token)
		username := claims[JWTUserNameKey].(string)

		ctx := context.WithValue(request.Context(), "username", username)
		endpointHandler(writer, request.WithContext(ctx))
	}
}

func GenerateJWTToken(username string) (string, error) {
	return jwt.CreateToken(JWTUserNameKey, username)
}
