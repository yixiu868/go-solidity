package main

import "github.com/yixiu868/go-solidity/pkg/gobase/task3/gorm"

func main() {
	//configPath := filepath.Join(".", "configs", "development.yaml") // 假设工作目录是项目根目录
	//// 加载配置文件
	//config, err := configs.LoadConfig(configPath)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = db.InitDB(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer db.CloseDB()
	//
	//commentRepo := repository.NewCommentRepository(db.DB)
	//postRepo := repository.NewPostRepository(db.DB)
	//userRepo := repository.NewUserRepository(db.DB)
	//
	//// 迁移schema
	//commentRepo.AutoMigrate()
	//postRepo.AutoMigrate()
	//userRepo.AutoMigrate()
	//
	//ctx := context.Background()

	// 新增user
	//email := "lisi@163.com"
	//birthday := time.Now().AddDate(-20, 0, 0)
	//memberNumber := "abc001"
	//user1 := model.User{
	//	Name:         "lisi",
	//	Email:        &email,
	//	Age:          20,
	//	Birthday:     &birthday,
	//	MemberNumber: sql.NullString{String: memberNumber, Valid: true},
	//	ActivatedAt: sql.NullTime{
	//		Time:  time.Now(),
	//		Valid: true,
	//	},
	//}
	//
	//userRepo.Create(ctx, &user1)

	// 新增post
	//post1 := model.Post{
	//	Title:  "任逍遥",
	//	Body:   "沧海一声笑",
	//	Author: "lisi",
	//	UserID: 2,
	//}
	//postRepo.Create(ctx, &post1)

	// 新增评论
	//comment1 := model.Comment{
	//	Content: "花儿为什么那么红，是一首好听的歌",
	//	PostID:  1,
	//}
	//commentRepo.Create(ctx, &comment1)
	//comment2 := model.Comment{
	//	Content: "花儿为什么那么红，哎呦不错哟",
	//	PostID:  1,
	//}
	//commentRepo.Create(ctx, &comment2)
	//comment3 := model.Comment{
	//	Content: "任逍遥，听出逍遥自在的感觉",
	//	PostID:  3,
	//}
	//commentRepo.Create(ctx, &comment3)

	gorm.AssociationSearch()
}
