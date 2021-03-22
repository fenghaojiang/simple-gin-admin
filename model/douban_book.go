package model

type DoubanBook struct {
	Id         int32   `db:"id" json:"id"`
	Title      string  `db:"title" json:"title"`
	Author     string  `db:"author" json:"author"`
	Translator string  `db:"translator" json:"translator"`
	Press      string  `db:"press" json:"press"`
	Date       string  `db:"date" json:"date"`
	Price      string  `db:"price" json:"price"`
	Star       float64 `db:"star" json:"star"`
	CommentNum int32   `db:"comment" json:"comment_num"`
	Quote      string  `db:"quote" json:"quote"`
}

func (d DoubanBook) TableName() string {
	return "douban_book"
}
