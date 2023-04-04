package main

import (
	"gin-vue/common"
	"gin-vue/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()          //加在配置类
	db := common.InitDB() //初始化数据库
	defer db.Close()
	r := gin.Default()
	r = route.CollectRoute(r)              //配置路由
	port := viper.GetString("server.port") //获取配置类所设置的端口号
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	//初始化配置类
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}
}
