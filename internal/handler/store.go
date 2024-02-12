package handler

import "github.com/khang00/verbose-spork/internal/model"

type UserStore interface {
	CreateUser(username string, password string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
}

type KeywordStore interface {
	CreateKeywords(keywords []*model.Keyword) ([]*model.Keyword, error)
	GetKeywordByID(ID uint) (*model.Keyword, error)
}
