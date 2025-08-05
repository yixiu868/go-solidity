package main

import (
	"github.com/yixiu868/go-solidity/configs"
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

//func main() {
//	posts, err := repo.GetUserPostsWithAllComments(db.DB, 2)
//	if err != nil {
//		panic(err)
//	}
//	postsJSON, err := json.Marshal(posts)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(postsJSON))
//
//	posts, err = repo.GetMostCommentPost(db.DB)
//	if err != nil {
//		panic(err)
//	}
//	postsJSON, err = json.Marshal(posts)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(postsJSON))
//}
