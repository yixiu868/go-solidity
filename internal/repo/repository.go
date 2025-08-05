package repo

import (
	"github.com/yixiu868/go-solidity/internal/model"
	"gorm.io/gorm"
)

// 迁移数据表
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}

// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func GetUserPostsWithAllComments(db *gorm.DB, userID uint) ([]model.Post, error) {
	var Posts []model.Post
	return Posts, db.Where("user_id = ?", userID).Preload("Comments").Find(&Posts).Error
}

// 查询评论数量最多的文章信息
func GetMostCommentPost(db *gorm.DB) ([]model.Post, error) {
	var Posts []model.Post
	result := db.Model(&model.Post{}).
		Select("post.*, COUNT(comment.id) as comment_count").
		Joins("LEFT JOIN comment ON post.id = comment.post_id").
		Group("post.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&Posts)
	return Posts, result.Error
}
