package main

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/dao"
	"github.com/3th-JustFunTeam/Tiktok-Backend/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化grom
	dao.InitDB()

	r := gin.Default()
	server.InitRouter(r)
	r.Run()
}
