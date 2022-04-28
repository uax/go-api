package models

type Book struct {
	Id          uint   `json:"id"`
	UserId      int    `json:"user_id"`      //所属用户 ID
	Channel     int8   `json:"channel"`      //男女频标识，男频0，女频1
	Name        string `json:"name"`         //书籍名称
	Thumb       string `json:"thumb"`        //封面
	Author      string `json:"author"`       //作者
	Chapters    uint   `json:"chapters"`     //章节数
	Length      int    `json:"length"`       //小说字数
	CategoryId  int    `json:"category_id"`  //分类 ID
	WriteStatus int8   `json:"write_status"` //连载/完结:1/0
	Description string `json:"description"`  //简介
	Tag         string `json:"tag"`          //标签
	Last        string `json:"last"`         //最后更新章节
	Status      int8   `json:"status"`       // 上架/下架:1/0
	Topic       int8   `json:"topic"`        //是否推荐
	Sign        int8   `json:"sign"`         //是否签约
	Publish     int8   `json:"publish"`      //是否发布到API
	Check       int8   `json:"check"`        //章节是否需要审核：1/审核,0/无需审核
	Visit       int64  `json:"visit"`        //阅读量
	Free        int8   `json:"free"`         //是否免费:1/免费，0/收费
	Origin      int8   `json:"origin"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func (book Book) List(id string) string {
	return "res" + id
}
