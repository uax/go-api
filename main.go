package main

import (
	"api/config"
	"api/db"

	"api/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	// 初始化数据库连接池
	if err := db.InitDb(); err != nil {
		log.Fatal("初始化错误:", err)
	}
}

func main() {
	gin.SetMode(config.Cfg.Section("").Key("app_mode").String())
	router := router.InitRouter()
	router.Run(config.Cfg.Section("server").Key("host_name").String() + ":" + config.Cfg.Section("server").Key("host_port").String())
}
