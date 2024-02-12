package jwt

import (
	"github.com/golang-jwt/jwt"
	"testing"
)

func TestParseJWT(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "create token",
			args: args{username: "test"},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QifQ.FILrByQNl1Mx00RSZonmT3p5waGlFaymZJ3e3a5VBac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken("username", tt.args.username)
			if err != nil {
				t.Errorf("CreateToken() error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	type args struct {
		jwtString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "validate token",
			args: args{jwtString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QyIn0.qONKoovV3NaLnFhZl5wO0Leh6gEAF09e13fqPpN6aiI"},
			want: "test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateToken(tt.args.jwtString)
			if err != nil {
				t.Errorf("ValidateToken() error = %v", err)
				return
			}

			claims, ok := got.Claims.(jwt.MapClaims)
			if !ok || !got.Valid {
				t.Errorf("ValidateToken() token is not valid got = %v", got)
			}

			username := claims["username"].(string)
			if username != tt.want {
				t.Errorf("ValidateToken() token's claims 'username' is wrong got = %s want = %s", username, tt.want)
			}
		})
	}
}
