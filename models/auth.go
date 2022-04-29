package models

import "time"

type Auth struct {
	Id        uint64    `gorm:"primary_key;type:bigint(20) auto_increment;not null;comment:'ID';" json:"id"`
	Openid    string    `gorm:"size:32;notnull;index:idx_openid,unique comment:'Openid';" json:"openid"`
	Name      string    `gorm:"size:255;" json:"name"`
	Avatar    string    `gorm:"size:255;" json:"avatar""`
	CreatedAt time.Time `json:"created_at"`
}
