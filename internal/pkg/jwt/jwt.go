package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

var defaultSecretKey = []byte("secret")

func CreateToken(payload map[string]string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for k, v := range payload {
		claims[k] = v
	}

	return token.SignedString(defaultSecretKey)
}

func ValidateToken(jwtString string) (*jwt.Token, error) {
	return jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}

		return defaultSecretKey, nil
	})
}

func GetClaims(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}
