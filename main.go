package main

import (
	. "api/config"
	"api/db"
	//_ "api/database"
	"api/middleware"
	"api/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	// 初始化配置文件
	err := LoadConfig()
	if err != nil {
		log.Fatal("初始化错误:", err)
	}

	// 初始化数据库连接池
	if err = db.InitDb(ConfAll.DB); err != nil {
		log.Fatal("初始化错误:", err)
	}
}

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
