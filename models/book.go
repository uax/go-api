package models

type Book struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"thumb"`
}

func (book Book) List(id string) string {
	return "res" + id
}
