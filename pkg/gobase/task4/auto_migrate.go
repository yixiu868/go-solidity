package main

import (
	"github.com/yixiu868/go-solidity/configs"
	"github.com/yixiu868/go-solidity/internal/model/blog"
	"github.com/yixiu868/go-solidity/pkg/gobase/db"
	"path/filepath"
)

func init() {
	configPath := filepath.Join(".", "configs", "development.yaml") // 假设工作目录是项目根目录
	// 加载配置文件
	config2, err := configs.LoadConfig2(configPath)
	if err != nil {
		panic(err)
	}

	err = db.InitDB2(config2)
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	db.DB2.AutoMigrate(&blog.User{}, &blog.Post{}, &blog.Comment{})
	defer db.CloseDB2()
}

func main() {
	Migrate()
}
