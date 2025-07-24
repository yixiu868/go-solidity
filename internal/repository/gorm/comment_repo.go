package gorm

import (
	"context"
	gorm2 "github.com/yixiu868/go-solidity/internal/model/gorm"
	"gorm.io/gorm"
)

type CommentRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, comment *gorm2.Comment) error
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
	return r.db.AutoMigrate(&gorm2.Comment{})
}

// 新增
func (r *commentRepository) Create(ctx context.Context, comment *gorm2.Comment) error {
	return r.db.Create(comment).Error
}
