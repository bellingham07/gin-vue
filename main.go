package main

import (
	"gin-vue/common"
	"gin-vue/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = route.CollectRoute(r)
	panic(r.Run())
}
