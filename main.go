package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"simple-gin-admin/conf"
	"simple-gin-admin/controller"
	"simple-gin-admin/dao"
	"simple-gin-admin/middleware"
)

func main() {
	err := conf.Parse()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Fail to load config.toml\n")
	}
	dao.NewDao()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/", controller.GetSomeData)
	other := r.Group("/other")
	{
		other.POST("/otherData", controller.GetSomeOtherData)
	}
	err = r.Run(conf.GlobalConfig.Port)
	if err != nil {
		log.Fatal("Fail to Run gin")
	}
}
