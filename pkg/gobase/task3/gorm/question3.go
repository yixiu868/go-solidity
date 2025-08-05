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
	//db.DB.Create(&model.Post{Title: "呐喊", Content: "呐喊》收录作者1918年至1922年所作小说十四篇。1923年8月由北京新潮社出版，原收十五篇，列为该社《文艺丛书》之一。1924年5月第三次印刷时起，改由北京北新书局出版，列为作者所编的《乌合丛书》之一。1930年1 月第十三次印刷时，由作者抽去其中的《不周山》一篇(后改名为《补天》，收入《故事新编》)。作者生前共印行二十二版次。", UserID: 1})
	// 增加评论
	//db.DB.Create(&model.Comment{Content: "修行129th，小说写到这份上，已经无话可说，读每一篇都要击节大赞一个好字，最喜欢的，反而是最写实，最无弦外之音的社戏。另外，大先生小说拿底层人开刀的频率真是略高啊，那时代的底层人无力阅读，也无力还手，讽刺他们又有何意义呢，名垂千古的阿Q正传，空前绝后的好，但是读来，确实有点不舒服，这个问题也不是两三百字说的清楚的，有空码长文给自己看吧", UserID: 1, PostID: 5})
	db.DB.Delete(&model.Comment{ID: 7, PostID: 5})
}
