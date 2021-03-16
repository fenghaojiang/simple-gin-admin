package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"simple-gin-admin/conf"
	"simple-gin-admin/controller"
	"simple-gin-admin/middleware"
)

func main() {
	err := conf.Parse()
	if err != nil || conf.GlobalConfig == nil {
		log.Fatal("Fail to load config.toml\n")
	}
	r := gin.Default()
	err = r.Run(conf.GlobalConfig.Port)
	if err != nil {
		log.Fatal("Fail to Run gin")
	}
	r.Use(middleware.CORSMiddleware())
	r.GET("/", controller.Login)
	other := r.Group("/other")
	{
		other.POST("/otherPost", controller.Post)
	}
}
