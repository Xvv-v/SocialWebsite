package repository

import (
	"log"

	"github.com/jinzhu/gorm"
)

//连接mysql数据库

var (
	mysqldb *gorm.DB
)

//InitMysql 初始化
func InitMysql() *gorm.DB {

	var err error
	mysqldb, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/SocialWebsite")
	if err != nil {
		log.Fatal(err)
	}
	return mysqldb
}

//CloseDB 关闭
func CloseDB() {
	mysqldb.Close()
}
