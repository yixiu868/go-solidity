package gorm

import (
	"context"
	gorm2 "github.com/yixiu868/go-solidity/internal/model/gorm"
	"gorm.io/gorm"
)

type PostRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, post *gorm2.Post) error
	FindByUserId(ctx context.Context, userId uint) ([]gorm2.Post, error)
	FindMostCommentPost(ctx context.Context) (*gorm2.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

// 迁移schema
func (r *postRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&gorm2.Post{})
}

// 新增
func (r *postRepository) Create(ctx context.Context, post *gorm2.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindByUserId(ctx context.Context, userId uint) ([]gorm2.Post, error) {
	var posts []gorm2.Post
	err := r.db.Model(&gorm2.Post{}).Preload("Comments").Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}

// 评论数最多的帖子
func (r *postRepository) FindMostCommentPost(ctx context.Context) (*gorm2.Post, error) {
	var post gorm2.Post
	tx := r.db.Model(&gorm2.Post{}).Where("id = (?)", r.db.Model(&gorm2.Comment{}).Select("post_id").Group("post_id").Order("count(1) desc").Limit(1)).First(&post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &post, nil
}
