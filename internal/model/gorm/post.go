package gorm

import (
	"fmt"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string
	Body     string
	Author   string
	UserID   uint
	Comments []Comment
}

func (post *Post) AfterSave(db *gorm.DB) error {
	var user User
	// 直接通过 post.UserID 查询用户，避免子查询
	tx := db.Preload("Posts").First(&user, post.UserID)
	if tx.Error != nil {
		return tx.Error
	}

	fmt.Println("作者", user.Name, "文章数达", len(user.Posts), "篇")
	return nil
}
