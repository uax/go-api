package models

import "time"

type Chapter struct {
	Id          uint      `json:"id"`
	UserId      int       `json:"user_id"` //所属用户 ID
	BookId      int       `json:"book_id"` //所属书籍 ID
	Title       string    `json:"title"`   //章节标题
	Content     string    `json:"content"` //章节内容
	Source      string    `json:"source"`
	Draft       int8      `json:"draft"`  //草稿:否/0,是/1
	Status      int8      `json:"status"` //待审/正常/拒绝:0/1/2
	Vip         int8      `json:"vip"`    //VIP 是/否:1/0
	Price       int       `json:"price"`  //价格:分
	Order       int       `json:"order"`  //章节排序
	Length      int       `json:"length"` //章节字数
	CollectId   int       `json:"collect_id"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
}

// TableName 会将 Chapter 的表名重写为 `sections`
func (Chapter) TableName() string {
	return "sections"
}
