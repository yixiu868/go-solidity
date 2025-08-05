package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Password  string         `gorm:"size:50;not null" json:"-"`
	Email     string         `gorm:"size100;" json:"email"`
	PostCount int            `gorm:"default:0" json:"post_count"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Posts     []Post         `gorm:"foreignKey:UserID" json:"posts,omitempty"` // 关联用户的文章
}

func (u *User) TableName() string {
	return "user"
}

type Post struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title         string         `gorm:"size:200;not null" json:"title"`
	Content       string         `gorm:"type:text;not null" json:"content"`
	UserID        uint           `gorm:"not null;index" json:"user_id"`           // 关联用户的ID
	CommentStatus string         `gorm:"size:20;default:'无评论'"`                   // 评论状态："有评论"或"无评论"
	User          User           `gorm:"foreignKey:UserID" json:"user,omitempty"` // 关联的用户信息
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-"`
	Comments      []Comment      `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	CommentCount  int            `gorm:"-" json:"comment_count"` // 非数据库字段，用于存储评论数量
}

func (a *Post) TableName() string {
	return "post"
}

func (a *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).
		Where("id = ?", a.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

type Comment struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string         `gorm:"size:500;not null" json:"content"`
	PostID    uint           `gorm:"not null;index" json:"post_id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Post      Post           `gorm:"foreignKey:PostID" json:"post,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 1. 查询当前文章的剩余有效评论数（排除已软删除的评论）
	var commentCount int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ? AND deleted_at IS NULL", c.PostID). // 只统计未删除的评论
		Count(&commentCount).Error; err != nil {
		return err // 若查询失败，返回错误
	}

	// 2. 若评论数为0，更新文章的评论状态
	if commentCount == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").
			Error
	}

	return nil
}
