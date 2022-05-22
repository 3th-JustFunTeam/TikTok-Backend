package main

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化grom并连接mysql
	dao.InitDB()

	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
