package repository

import (
	"context"
	"github.com/yixiu868/go-solidity/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, post *model.Post) error
	FindByUserId(ctx context.Context, userId uint) ([]model.Post, error)
	FindMostCommentPost(ctx context.Context) (*model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

// 迁移schema
func (r *postRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&model.Post{})
}

// 新增
func (r *postRepository) Create(ctx context.Context, post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindByUserId(ctx context.Context, userId uint) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Model(&model.Post{}).Preload("Comments").Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindMostCommentPost(ctx context.Context) (*model.Post, error) {
	return nil, nil
}
