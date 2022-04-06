package main

import (
	"api/base"
	"api/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORS)
	router.GET("/", func(ctx *gin.Context) {
		base.WrapContext(ctx).Success("hello,world")
	})
	router.GET("/register", func(ctx *gin.Context) {
		jwtAuth := &middleware.JwtAuth{SignKey: []byte("noah")}
		tokenStr, err := jwtAuth.GenerateToken(3600, "noah")
		if err != nil {
			base.WrapContext(ctx).Error(401, "GenerateToken Error")
		}
		base.WrapContext(ctx).Success(tokenStr)
		//base.WrapContext(ctx).Success("hello,world")
	})
	user := router.Group("/user").Use(middleware.JWT)
	{
		user.GET("/info", func(ctx *gin.Context) {
			claims, exist := ctx.Get(middleware.ClaimsKey)
			if !exist {
				fmt.Println("获取用户信息失败")
			}
			fmt.Printf("SUCCESS  %s", claims)
		})
	}
	router.Run(":8000")
}
