package models

import "api/database"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

var Users []User

//Insert insert user
func (user User) Insert() (id int64, err error) {
	result := database.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//Users get users
func (user *User) Users() (users []User, err error) {
	if err = database.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}
