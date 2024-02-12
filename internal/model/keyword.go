package model

import "gorm.io/gorm"

type Keyword struct {
	gorm.Model
	Keyword       string
	ResultStats   int
	NumberOfLinks int
	NumberOfAds   int
	HTML          string

	UserID uint
	User   User
}
