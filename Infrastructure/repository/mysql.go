package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//Mysqldb mysql数据库链接
	Mysqldb *gorm.DB
)

//初始化
func init() {

	var err error
	Mysqldb, err = gorm.Open("mysql", "root:jingtt106@tcp(127.0.0.1:3306)/SocialWebsite?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
}
