package main

import (
	"fmt"
	"github.com/3th-JustFunTeam/Tiktok-Backend/config"
	"github.com/3th-JustFunTeam/Tiktok-Backend/server"
	"github.com/gin-gonic/gin"
)

func main() {
	var config config.Config
	//config.ReadConfig("./config/config.production.yml")
	config.ReadConfig("./config.development.yml")
	fmt.Printf("%v\n", config.MySQL)
	r := gin.Default()
	server.InitRouter(r)
	r.Run()
}
