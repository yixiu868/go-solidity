package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yixiu868/go-solidity/configs"
	"github.com/yixiu868/go-solidity/internal/model/blog"
	"github.com/yixiu868/go-solidity/pkg/gobase/db"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"path/filepath"
	"time"
)

var JWTKey = []byte("oIft1b5qZjyLcc0zZo2UrUx5rk3KE0LvZKv73fw502oXd6vfYu1OAQvbSel8whvm")

func main() {
	// 初始化数据库连接池
	configPath := filepath.Join(".", "configs", "development.yaml") // 假设工作目录是项目根目录
	// 加载配置文件
	config2, err := configs.LoadConfig2(configPath)
	if err != nil {
		panic(err)
	}
	db.InitDB2(config2)

	// 程序退出关闭数据库连接池
	defer db.CloseDB2()

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

	// 创建文章
	router.POST("/createArticle", func(c *gin.Context) {})

	// 更新文章
	router.POST("/updateArticle/", func(c *gin.Context) {})

	// 删除文章
	router.POST("/deleteArticle/", func(c *gin.Context) {})

	// 创建评论

	//

	router.Run()
}

// 注册
func Register(c *gin.Context) {
	var user blog.User
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

	if err := db.DB2.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user blog.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser blog.User
	if err := db.DB2.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.Header("x-jwt-token", tokenStr)
	c.JSON(http.StatusOK, gin.H{
		"success": "Login successfully",
	})
}
