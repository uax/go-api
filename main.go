package main

import (
	//_ "api/database"
	"api/middleware"
	"api/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := router.InitRouter()
	router.Run(":8000")
}

type User struct {
	Name string `json:"name"`
}

func authHandle(c *gin.Context) {
	user := User{}
	err := c.BindJSON(&user) // c.ShouldBind(&user)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "无效的参数",
		})
		return
	}
	fmt.Println(user)
	if user.Name == "noah" {
		tokenString, _ := middleware.GenerateToken(381479)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "ok",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"msg":  "登录失败",
	})
	return
}
