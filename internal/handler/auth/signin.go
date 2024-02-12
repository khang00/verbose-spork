package auth

import (
	"encoding/json"
	"fmt"
	"github.com/khang00/verbose-spork/internal/handler"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninResponse struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func (s *AuthHandler) Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	req := &SigninRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.signin(req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *AuthHandler) signin(req *SigninRequest) (*SigninResponse, error) {
	user, err := s.userStore.FindUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	token, err := handler.GenerateJWTToken(user.Username, strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, err
	}

	return &SigninResponse{
		UserID: int(user.ID),
		Token:  token,
	}, nil
}
