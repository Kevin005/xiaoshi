package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}

type Feedbacks struct {
	Model
	UserId  int    `json:"user_id"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

type Users struct {
	Model
	CreateTime  int64 `json:"create_time"`
	UserName    string    `json:"user_name"`
	NickName    string    `json:"nick_name"`
	PhoneNumber string    `json:"phone_number"`
	Age         int       `json:"age"`
	Gender      int       `json:"gender"`
	Token       string    `json:"token"`
}

type Books struct {
	Model
	CreateTime int64 `json:"create_time"`
	BookName   string    `json:"book_name"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	Image      string    `json:"image"`
	PageTotal  int       `json:"page_total"`
}

type MyBooks struct {
	Model
	UserId   uint   `json:"user_id"`
	BookId   uint   `json:"book_id"`
	Progress string `json:"progress"`
	Private  bool   `json:"private"`
}

type Posts struct {
	Model
	PostTime int64 `json:"post_time"`
	UserId   uint      `json:"user_id"`
	MyBookId uint      `json:"my_book_id"`
	Title    string    `json:"title"`
	Content  string    `json:content`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Feedbacks{}, &Users{}, &MyBooks{}, &Books{}, &Posts{})
	return db
}
