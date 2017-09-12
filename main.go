package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"cms/config"
	"cms/framework"
	_ "cms/database/mysql"
	"log"
	"fmt"
)

func main() {

	if config.AppConfig.Server.LogModelEnable {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	router.Use(cors.New(corsConfig))
	framework.LoadHandler(router)

	addr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	if gin.Mode() == gin.ReleaseMode {
		log.Printf("Listening and serving HTTP on %s", addr)
	}
	router.Run(addr)

}


