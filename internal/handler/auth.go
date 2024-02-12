package handler

import (
	"encoding/json"
	"fmt"
	"github.com/khang00/verbose-spork/internal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserStore interface {
	CreateUser(username string, password string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
}

type AuthHandler struct {
	userStore UserStore
}

func NewAuthHandler(userStore UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
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

	token, err := GenerateJWTToken(user.Username, strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, err
	}

	return &SignupResponse{
		UserID: int(user.ID),
		Token:  token,
	}, nil
}

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

	token, err := GenerateJWTToken(user.Username, strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, err
	}

	return &SigninResponse{
		UserID: int(user.ID),
		Token:  token,
	}, nil
}
