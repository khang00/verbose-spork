package store

import "github.com/khang00/verbose-spork/internal/model"

type Store interface {
	CreateUser(username string, password string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)

	CreateKeywords(keywords []*model.Keyword) ([]*model.Keyword, error)
}
