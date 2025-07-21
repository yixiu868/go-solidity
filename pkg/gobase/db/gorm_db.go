package db

import (
	"fmt"
	"github.com/yixiu868/go-solidity/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB(cfg *configs.Config) error {
	if DB != nil {
		return nil
	}

	// 构建DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		cfg.Database.MySQL.Username,
		cfg.Database.MySQL.Password,
		cfg.Database.MySQL.Host,
		cfg.Database.MySQL.Port,
		cfg.Database.MySQL.DBName,
		cfg.Database.MySQL.Charset,
		cfg.Database.MySQL.ParseTime,
		cfg.Database.MySQL.Loc,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 设置连接池参数
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return nil
}

func CloseDB() {
	if DB != nil {
		db, err := DB.DB()
		if err == nil {
			db.Close()
		}
		DB = nil
	}
}
