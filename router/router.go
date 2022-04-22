package router

import (
	"api/controllers"
	"api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	router.GET("/books/:id", controllers.BookList)

	auth := router.Group("/auth").Use(middleware.JWTMiddleware())
	{
		auth.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"code": 200, "msg": "hello,api"})
		})
		auth.GET("home", controllers.Home)
		//auth.GET("/auth/info", middleware.JWTMiddleware(), getInfo)
	}
	router.POST("/auth/login", controllers.Login)
	router.GET("users", controllers.Users)
	return router
	//router.Run(":8000")
}
