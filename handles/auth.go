package handles

import (
	. "api/config"
	"api/db"
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
	err := c.BindJSON(&json)
	if err != nil {
		fmt.Println("get code error.")
		return
	}
	//根据 code 获取 openid
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     Cfg.Section("miniapp").Key("appid").String(),
		AppSecret: Cfg.Section("miniapp").Key("appsecret").String(),
		Cache:     memory,
	}
	auth := wc.GetMiniProgram(cfg).GetAuth()
	result, err := auth.Code2Session(json.Code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	if result.ErrCode != 0 {
		fmt.Println("error" + result.ErrMsg)
		return
	}
	var user models.User
	res := db.ORM.Where("openid = ?", result.OpenID).First(&user)
	if res.RowsAffected > 0 {
		tokenString, _ := middleware.GenerateToken(user.ID)
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "ok",
			"token": tokenString,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  res.Error.Error(),
	})
	return
}
