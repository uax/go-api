package handles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Home getHome
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "home",
	})
}
