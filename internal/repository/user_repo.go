package repository

import (
	"context"
	"github.com/yixiu868/go-solidity/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// 迁移schema
func (r *userRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&model.User{})
}

// 新增
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.Create(user).Error
}
