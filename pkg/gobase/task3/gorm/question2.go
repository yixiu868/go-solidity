package gorm

import (
	"context"
	"fmt"
	"github.com/yixiu868/go-solidity/configs"
	"github.com/yixiu868/go-solidity/internal/repository"
	"github.com/yixiu868/go-solidity/pkg/gobase/db"
	"path/filepath"
)

func init() {
	configPath := filepath.Join(".", "configs", "development.yaml") // 假设工作目录是项目根目录
	// 加载配置文件
	config, err := configs.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	err = db.InitDB(config)
	if err != nil {
		panic(err)
	}
}

func AssociationSearch() {
	post_repo := repository.NewPostRepository(db.DB)
	posts, err := post_repo.FindByUserId(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(posts)

	defer db.CloseDB()
}
