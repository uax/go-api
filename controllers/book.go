package controllers

import (
	DB "api/database"
	model "api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BookList(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Param("id")) // c.Param("id")
	res := DB.Eloquent.Debug().Where("category_id = ? AND status = 1", bookId).Order("id desc").Find(&[]model.Book{})
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  res.Error.Error(),
			"data": []model.Book{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res.Value,
	})
}
