package main

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
	////
	//commentRepo := gorm2.NewCommentRepository(db.DB)
	//postRepo := repo.NewPostRepository(db.DB)
	//userRepo := repo.NewUserRepository(db.DB)
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

	//// 新增post
	//post1 := model.Post{
	//	Title:  "活着",
	//	Body:   "靠活着活着的作者",
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
	//comment3 := gorm.Comment{
	//	Content: "活着写的感人",
	//	PostID:  8,
	//}
	//commentRepo.Create(ctx, &comment3)
	//
	//db.DB.Model(&gorm.Comment{}).Delete(&comment3)

	//users, err := userRepo.FindInfoByUsername(context.Background(), 2)
	//if err != nil {
	//	panic(err)
	//}
	//users_bytes, err := json.Marshal(users)
	//fmt.Println(string(users_bytes))
	//
	//users2, err := userRepo.FindInfoByUsername(context.Background(), 1)
	//if err != nil {
	//	panic(err)
	//}
	//users_bytes2, err := json.Marshal(users2)
	//fmt.Println(string(users_bytes2))

	//post, err := postRepo.FindMostCommentPost(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//marshal, err := json.Marshal(post)
	//fmt.Println(string(marshal))
}
