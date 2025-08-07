package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yixiu868/go-solidity/configs"
	"github.com/yixiu868/go-solidity/internal/model"
	"github.com/yixiu868/go-solidity/pkg/gobase/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"time"
)

var JWTKey = []byte("oIft1b5qZjyLcc0zZo2UrUx5rk3KE0LvZKv73fw502oXd6vfYu1OAQvbSel8whvm")

// Claims 定义JWT的payload
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	// 初始化数据库连接池
	configPath := filepath.Join(".", "configs", "development.yaml") // 假设工作目录是项目根目录
	// 加载配置文件
	config2, err := configs.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	db.InitDB(config2)

	// 程序退出关闭数据库连接池
	defer db.CloseDB()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 注册
	router.POST("/register", Register)

	// 登录
	router.POST("/login", Login)

	protected := router.Group("/api")
	protected.Use(JWTMiddleware())
	{
		// 创建文章
		protected.POST("/createArticle", createArticle)

		// 更新文章
		protected.POST("/updateArticle", updateArticle)

		// 删除文章
		protected.POST("/deleteArticle", deleteArticle)

		// 创建评论
		protected.POST("/createComment", createComment)

		// 删除评论
		protected.POST("/deleteComment", deleteComment)
	}

	router.Run()
}

func createComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if comment.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "评论内容不能为空"})
	}
	// 校验待评论文章
	var existPost model.Post
	tx := db.DB.Where("id = ?", comment.PostID).First(&existPost)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tx.Error.Error()})
		return
	}
	comment.UserID = existPost.UserID
	db.DB.Create(&comment)
	c.JSON(http.StatusOK, gin.H{
		"message": "评论发表成功",
	})
}

func deleteComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 校验待评论文章
	var existComment model.Comment
	tx := db.DB.Where("id = ?", comment.ID).First(&existComment)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tx.Error.Error()})
		return
	}
	db.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{
		"message": "评论删除成功",
	})
}

func deleteArticle(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 校验待评论文章
	var existPost model.Post
	tx := db.DB.Where("id = ? and user_id = ?", post.ID, c.MustGet("userID").(uint)).First(&existPost)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tx.Error.Error()})
		return
	}
	db.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "文章删除成功",
	})
}

func updateArticle(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 先检查post.ID是否有效
	var existPost model.Post
	tx := db.DB.Where("id = ? and user_id = ?", post.ID, c.MustGet("userID").(uint)).First(&existPost)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tx.Error.Error()})
		return
	}

	if post.Title == "" && post.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标题或文章内容不能同时为空"})
	}
	post.UserID = c.MustGet("userID").(uint)
	ctx := context.Background()
	_, err := gorm.G[model.Post](db.DB).Where("id = ?", post.ID).Updates(ctx, model.Post{Title: post.Title, Content: post.Content})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功",
	})
}

func createArticle(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is empty"})
		return
	}
	if post.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content is empty"})
		return
	}
	post.UserID = c.MustGet("userID").(uint)
	db.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "文章发表成功",
	})
}

// 注册
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser model.User
	if err := db.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	clams := &Claims{
		UserID:   storedUser.ID,
		Username: storedUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "auth-service",
			Subject:   "user-token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)

	tokenStr, err := token.SignedString(JWTKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.Header("x-jwt-token", tokenStr)
	c.JSON(http.StatusOK, gin.H{
		"success": "Login successfully",
	})
}

// 验证JWT令牌
func validateToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

// JWT中间件，用于保护需要登录的路由
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从x-jwt-token头中获取token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 检查token格式是否正确
		var tokenString string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			c.Abort()
			return
		}

		// 验证token
		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
