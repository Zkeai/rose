package db

import (
	"database/sql"
	"fmt"
	"github.com/Zkeai/MuCoinPay/McPay-go/configs"
	"github.com/Zkeai/MuCoinPay/McPay-go/pkg/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.C.Mysql.User,
		configs.C.Mysql.Password,
		configs.C.Mysql.Host,
		configs.C.Mysql.Port,
		configs.C.Mysql.Dbname,
	)

	// 创建数据库连接
	var err error
	println(dsn)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		logger.Fatal("無法連接到MySQL: %v", err)
	}

	// 设置数据库连接的最大打开连接数、最大空闲连接数和连接最大生存时间
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(60 * time.Minute)

	err = DB.Ping()
	if err != nil {
		logger.Fatal("無法連接到MySQL: %v", err)
	}
}

// CloseMySQL 关闭MySQL数据库连接
func CloseMySQL() {
	if DB != nil {
		defer func(DB *sql.DB) {
			err := DB.Close()
			if err != nil {

			}
		}(DB)
	}
}
