package handles

import (
	"api/db"
	model "api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type APIChapter struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	Vip       bool      `json:"vip"`
	UpdatedAt time.Time `json:"updated_at"`
}

//ChapterList 获取小说章节列表
func ChapterList(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Query("book_id"))
	if err != nil || bookId <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "invalid param",
			"data": []string{},
		})
		return
	}
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

	var chapters []APIChapter
	if err := db.ORM.Model(&model.Chapter{}).Where("book_id = ? AND status = 1", bookId).Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order("id desc").Find(&chapters).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	var total int64
	db.ORM.Model(&model.Chapter{}).Where("book_id = ? AND status = 1", bookId).Count(&total)
	pages := int(total) / p.PageSize
	if int(total)%p.PageSize != 0 {
		pages++
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  chapters,
		"total": total,
		"pages": pages,
	})
}

//ChapterShow 根据 ID 获取章节内容
func ChapterShow(c *gin.Context) {
	chapterId, err := strconv.Atoi(c.Param("id"))
	if err != nil || chapterId <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "invalid param",
			"data": []string{},
		})
		return
	}
	var chapter model.Chapter
	res := db.ORM.Where("id = ?", chapterId).First(&chapter)

	if res.RowsAffected > 0 && chapter.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": chapter,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "invalid param",
		"data": []string{},
	})
}
