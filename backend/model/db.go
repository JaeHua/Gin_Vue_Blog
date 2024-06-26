package model

import (
	"backend/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// 全局变量
var DB *gorm.DB

var err error

// InitDB 初始化数据库
func InitDB() {
	db, err := gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName))
	if err != nil {
		fmt.Printf("连接数据库出错，请检查配置", err)
	}
	//禁用默认表名的复数形式
	db.SingularTable(true)
	//迁移数据库
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	//sets max connecting nums
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(10 * time.Second)

	if db == nil {
		println("fuck")
	}
	DB = db
}
func GetDB() *gorm.DB {
	return DB
}
