package repository

import (
	"context"
	"github.com/yixiu868/go-solidity/internal/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, comment *model.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

// 迁移schema
func (r *commentRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&model.Comment{})
}

// 新增
func (r *commentRepository) Create(ctx context.Context, comment *model.Comment) error {
	return r.db.Create(comment).Error
}
