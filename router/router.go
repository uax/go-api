package router

import (
	"api/db"
	"api/handles"
	"api/middleware"
	model "api/models"
	"fmt"
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
	router.GET("install", Install)
	return router
}

func Install(c *gin.Context) {
	if err := db.ORM.AutoMigrate(model.User{}); err != nil {
		fmt.Println(err)
	}
}
