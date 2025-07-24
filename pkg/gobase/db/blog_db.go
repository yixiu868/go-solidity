package db

import (
	"fmt"
	"github.com/yixiu868/go-solidity/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB2 *gorm.DB

func InitDB2(cfg *configs.Config2) error {
	if DB2 != nil {
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
	DB2, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	// 设置连接池参数
	sqlDB, err := DB2.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return nil
}

func CloseDB2() {
	if DB2 != nil {
		db, err := DB2.DB()
		if err == nil {
			db.Close()
		}
		DB2 = nil
	}
}
