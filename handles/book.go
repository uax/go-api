package handles

import (
	"api/db"
	model "api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//APIBook 小说结构体
type APIBook struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type APIShow struct {
	Book     APIBook      `json:"book"`
	Chapters []APIChapter `json:"chapters"`
}

//BookList 获取分类小说列表
func BookList(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Query("category_id"))

	var p db.Page
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
	if err := db.ORM.Where("category_id = ? AND status = 1", categoryId).Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order("id desc").Find(&books).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	var total int64
	db.ORM.Model(&model.Book{}).Where("category_id = ? AND status = 1", categoryId).Count(&total)
	pages := int(total) / p.PageSize
	if int(total)%p.PageSize != 0 {
		pages++
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  books,
		"total": total,
		"pages": pages,
	})
}

//BookShow 获取小说详情
func BookShow(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Param("id"))
	book := APIBook{}
	//chapters := []int{1, 2, 3}
	var chapters []APIChapter

	result := db.ORM.Debug().Model(&model.Book{}).Where("id = ?", bookId).First(&book)
	if result.RowsAffected > 0 {
		db.ORM.Model(&model.Chapter{}).Debug().Order("id desc").Where("book_id = ?", bookId).Limit(10).Find(&chapters)
		show := APIShow{
			Book:     book,
			Chapters: chapters,
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": show,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  result.Error.Error(),
	})
}
