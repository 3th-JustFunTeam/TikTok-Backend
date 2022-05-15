package main

import (
	"fmt"
	"github.com/3th-JustFunTeam/Tiktok-Backend/config"
	"github.com/3th-JustFunTeam/Tiktok-Backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	var config config.Config
	config.GetConfig()
	fmt.Printf("%v\n", config)
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
