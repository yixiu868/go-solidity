package main

import (
	"github.com/yixiu868/go-solidity/configs"
	"github.com/yixiu868/go-solidity/internal/model"
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

func main() {
	db.DB.Create(&model.Post{Title: "在细雨中呼喊", Content: "《在细雨中呼喊》是二十世纪九十年代的经典文本，是余华的第一部长篇力作。小说描述了一位江南少年的成长经历和心灵历程。《在细雨中呼喊 》的结构来自于对时间的感受，确切地说是对记忆中的时间的感受，叙述者 天马行空地在过去、现在和将来这三个时间维度里自由穿行，将忆记的 碎片 穿插、结集、拼嵌完整。", UserID: 2})
}
