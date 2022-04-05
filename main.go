package main

import (
	"api/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORS)
	router.GET("/", func(ctx *gin.Context) {
		fmt.Println("hello,api")
		ctx.JSON(200, gin.H{"code": 200, "msg": "hello,api"})
	})
	router.Run(":8000")
}
