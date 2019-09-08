package models

import (
	"fmt"
	"github.com/gaogep/EchoPlay/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	baseUrl := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	DBUrl := fmt.Sprintf(baseUrl,
		settings.DBConf["USER"],
		settings.DBConf["PASSWORD"],
		settings.DBConf["HOST"],
		settings.DBConf["DBNAME"])

	var err error
	db, err = gorm.Open(settings.DBConf["DBTYPE"], DBUrl)
	if err != nil {
		log.Fatal("Connect database error: ", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 自动迁移数据库
	db.AutoMigrate(&User{}, &Post{}, &Category{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	defer db.Close()
}
