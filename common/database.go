package common

import (
	"fmt"
	"gin-vue/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "Zxcvbnm565"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)

	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}

	//gorm自动创建数据表
	db.AutoMigrate(&model.User{})

	DB = db //不赋值会出现空指针错误
	return db
}

func GetDB() *gorm.DB {
	return DB
}
