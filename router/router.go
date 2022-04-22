package router

import (
	"api/controllers"
	"api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	auth := router.Group("/").Use(middleware.JWTMiddleware())
	{
		auth.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"code": 200, "msg": "hello,api"})
		})
		auth.GET("home", controllers.Home)
		auth.GET("user", controllers.GetInfo)
	}
	router.POST("/login", controllers.Login)
	router.GET("/books/:id", controllers.BookList)
	router.GET("users", controllers.Users)
	return router
}
