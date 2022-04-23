package controllers

import (
	DB "api/database"
	model "api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BookList(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))
	var p DB.Page
	if c.ShouldBindQuery(&p) != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "invalid params",
		})
		return
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	switch {
	case p.PageSize > 100:
		p.PageSize = 100
	case p.PageSize <= 0:
		p.PageSize = 10
	}

	var books []model.Book
	if err := DB.Eloquent.Where("category_id = ? AND status = 1", categoryId).Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order("id desc").Find(&books).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	var total int
	DB.Eloquent.Model(&model.Book{}).Where("category_id = ? AND status = 1", categoryId).Count(&total)
	pages := total / p.PageSize
	if total%p.PageSize != 0 {
		pages++
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  books,
		"total": total,
		"pages": pages,
	})
}
