package auth

import (
	"encoding/json"
	"github.com/khang00/verbose-spork/internal/handler"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (s *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	req := &SignupRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.signup(req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *AuthHandler) signup(req *SignupRequest) (*SignupResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.userStore.CreateUser(req.Username, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	token, err := handler.GenerateJWTToken(user.Username, strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, err
	}

	return &SignupResponse{
		UserID:   int(user.ID),
		Username: user.Username,
		Token:    token,
	}, nil
}
