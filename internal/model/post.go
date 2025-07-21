package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Body     string
	Author   string
	UserID   uint
	Comments []Comment
}
