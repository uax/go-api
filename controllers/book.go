package controllers

import (
	DB "api/database"
	model "api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookList(c *gin.Context) {

	fmt.Println(c.Param("id"))
	bookId := c.Param("id")
	//res := model.Book{}.List(bookId)
	res := DB.Eloquent.Debug().Where("book_id=?", bookId).Find(&model.Book{})
	// 获取全部匹配的记录
	//db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
	return
}
