package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	PostID  uint
}
