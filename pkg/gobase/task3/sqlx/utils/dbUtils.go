package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

// 初始化连接数据库
func InitDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True"
	Db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql err:%v\n", err)
		return
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)
	return
}
