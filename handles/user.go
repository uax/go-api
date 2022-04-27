package handles

import (
	model "api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Users(c *gin.Context) {
	var user model.User
	user.Name = "noah"

	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "未找到相关信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
	})
}

func GetInfo(c *gin.Context) {
	uid := c.MustGet("uid")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"uid":  uid,
	})
}
