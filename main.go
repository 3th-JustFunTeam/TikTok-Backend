package main

import (
	"fmt"
	"github.com/3th-JustFunTeam/Tiktok-Backend/config"
	"github.com/3th-JustFunTeam/Tiktok-Backend/router"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// read configuration
	var c config.Config
	err := c.GetConfig()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v\n", c)

	// init mysql db
	config.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Asia%%2fShanghai", c.MySQL.User, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.MySQL.Database),
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// init redis db
	config.RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: "",
		DB:       0,
	})

	// init gin and routers
	r := gin.Default()
	router.InitRouter(r)

	// run app
	err = r.Run()
	if err != nil {
		panic(err.Error())
	}
}
