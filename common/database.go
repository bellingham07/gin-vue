package common

import (
	"fmt"
	"gin-vue/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)

	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}

	//gorm根据模型 自动创建数据表
	db.AutoMigrate(&model.User{})

	DB = db //不赋值会出现空指针错误
	return db
}

func GetDB() *gorm.DB {
	return DB
}
