package models

import (
	"api/db"
)

type User struct {
	Id              uint64 `json:"id"`
	Name            string `json:"name"`
	Openid          string `json:"openid"`
	Avatar          string `json:"avatar"`
	Email           string `json:"email"`
	EmailVerifiedAt int64  `json:"email_verified_at"`
	Password        string `json:"password"`
	RememberToken   string `json:"remember_token"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
}

var Users []User

//Insert insert user
func (user User) Insert() (id uint64, err error) {
	result := db.ORM.Create(&user)
	id = user.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//Users get users
func (user *User) Users() (users []User, err error) {
	if err = db.ORM.Find(&users).Error; err != nil {
		return
	}
	return
}

func (User) UserByOpenID(openid string) (u User, err error) {
	db.ORM.Debug().Where("openid = ?", openid).First(&u)
	//database.Eloquent.Debug().Where("openid = ?", openid).First(&u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
