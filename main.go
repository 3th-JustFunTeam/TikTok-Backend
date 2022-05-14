package main

import (
	"fmt"
	"github.com/3th-JustFunTeam/Tiktok-Backend/config"
	"github.com/3th-JustFunTeam/Tiktok-Backend/server"
	"github.com/gin-gonic/gin"
)

func main() {
	var config config.Config
	config.GetConfig()
	fmt.Printf("%v\n", config.MySQL)
	return
	r := gin.Default()
	server.InitRouter(r)
	r.Run()
}
