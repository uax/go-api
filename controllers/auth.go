package controllers

import (
	"api/middleware"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"net/http"
)

type Credential struct {
	Code string `json:"code"`
}

func Login(c *gin.Context) {
	var json Credential
	c.BindJSON(&json)
	//根据 code 获取 openid
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     "wx79ceafb3aa8bd583",
		AppSecret: "b825391c825400c531716c74dcd7110e",
		Cache:     memory,
	}
	auth := wc.GetMiniProgram(cfg).GetAuth()
	result, _ := auth.Code2Session(json.Code)
	if result.ErrCode != 0 {
		fmt.Println("error" + result.ErrMsg)
	}
	var user models.User
	user, _ = models.User.UserByOpenID(user, result.OpenID)
	tokenString, _ := middleware.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"code":  2000,
		"msg":   "ok",
		"token": tokenString,
	})
}
