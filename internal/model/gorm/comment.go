package gorm

import (
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	PostID  uint
}

// AfterDelete 评论删除后触发的钩子
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var commentCount int64
	// 注意：必须用 tx 而非全局 db，确保事务一致性（若删除操作在事务中）
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err // 统计失败时返回错误
	}

	// 2. 若评论数量为 0，更新对应文章的 CommentStatus 为“无评论”
	if commentCount == 0 {
		fmt.Println("文章id", c.PostID, "已无评论")
	}

	return nil
}
