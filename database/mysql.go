package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {

	var err error
	Eloquent, err = gorm.Open("mysql", "root:526937@tcp(127.0.0.1:3306)/api?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	if err != nil {
		panic("failed to content database")
	}
	if Eloquent.Error != nil {
		fmt.Printf("database error: %v", Eloquent.Error)
	}
}
