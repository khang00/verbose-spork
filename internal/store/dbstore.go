package store

import (
	"fmt"
	"github.com/khang00/verbose-spork/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultConnectionString = "host=localhost user=postgres password=12345678 dbname=postgres port=5432 sslmode=disable"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

type PostgresStore struct {
	conn *gorm.DB
}

func NewPostgresStore(config *PostgresConfig) (Store, error) {
	dsn := defaultConnectionString
	if config != nil {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.Host, config.User, config.Password, config.DBName, config.Port)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Keyword{})
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		conn: db,
	}, nil
}

func (s *PostgresStore) CreateUser(username string, password string) (*model.User, error) {
	user := &model.User{Username: username, Password: password}

	err := s.conn.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := s.conn.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) CreateKeywords(keywords []*model.Keyword) ([]*model.Keyword, error) {
	err := s.conn.Create(&keywords).Error
	if err != nil {
		return nil, err
	}

	return keywords, nil
}
