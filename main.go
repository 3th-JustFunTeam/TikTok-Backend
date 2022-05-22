package main

import (
	"github.com/3th-JustFunTeam/Tiktok-Backend/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server.InitRouter(r)
	r.Run()
}
