package router

import (
	"api/handles"
	"api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	auth := router.Group("/").Use(middleware.JWTMiddleware())
	{
		auth.GET("/", handles.Home)
		auth.GET("/home", handles.Home)
		auth.GET("user", handles.GetInfo)
	}
	router.POST("/login", handles.Login)
	router.GET("/category/:id", handles.Category)
	router.GET("/books/:id", handles.Show)
	router.GET("users", handles.Users)
	return router
}
