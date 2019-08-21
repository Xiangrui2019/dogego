package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase(connectionDSN string) {
	db, err := gorm.Open("mysql", connectionDSN)

	if err != nil {
		log.Fatal(err)
	}

	// 连接池设置
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	err = MigrationModels()

	if err != nil {
		log.Fatal(err)
	}
}
