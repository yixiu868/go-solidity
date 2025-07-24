package gorm

import (
	"context"
	gorm2 "github.com/yixiu868/go-solidity/internal/model/gorm"
	"gorm.io/gorm"
)

type UserRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, user *gorm2.User) error
	FindInfoByUsername(ctx context.Context, id int) (*gorm2.User, error)
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
	return r.db.AutoMigrate(&gorm2.User{})
}

// 新增
func (r *userRepository) Create(ctx context.Context, user *gorm2.User) error {
	return r.db.Create(user).Error
}

// 查询某个用户发布的所有文章及其对应的评论信息
func (r *userRepository) FindInfoByUsername(ctx context.Context, id int) (*gorm2.User, error) {
	var user gorm2.User
	tx := r.db.Where("id = ?", id).Preload("Posts.Comments").Preload("Posts").Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
