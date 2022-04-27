package db

import (
	. "api/config"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Url            = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s"
	MaxOpen        = 5 // 最大打开数
	MaxIdle        = 2 // 最大保留连接数
	LifeMinuteTime = 5 // 连接可重用最大时间
)

var ORM *gorm.DB

func InitDb() error {
	var err error
	var sqlDb *sql.DB
	ORM, err = gorm.Open(mysql.Open(
		fmt.Sprintf(
			Url, Cfg.Section("db").Key("username").String(), Cfg.Section("db").Key("password").String(),
			Cfg.Section("db").Key("host").String(), Cfg.Section("db").Key("port").MustInt(3306), Cfg.Section("db").Key("database").String())), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDb, err = ORM.DB()

	if err != nil {
		return err
	}
	sqlDb.SetMaxOpenConns(MaxOpen)
	sqlDb.SetConnMaxIdleTime(MaxIdle)
	sqlDb.SetConnMaxLifetime(LifeMinuteTime * time.Minute)
	return nil
}

type Page struct {
	PageNum  int    `form:"page"`
	PageSize int    `form:"size"`
	Keyword  string `form:"keyword"`
	Desc     string `form:"desc"`
}
