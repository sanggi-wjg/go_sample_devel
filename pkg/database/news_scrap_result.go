package database

import (
	"time"
)

type NewsScrapResult struct {
	Id    uint64 `gorm:"primaryKey;autoIncrement"`
	Href  string `gorm:"column:href;size:500;notnull;"`
	Title string `gorm:"column:title;index;size:500;notnull;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (NewsScrapResult) TableName() string {
	return "news_scrap_results"
}

//func CreateNewsScrapResult(href, title string) (*NewsScrapResult, error) {
//	news := NewsScrapResult{Href: href, Title: title}
//	res := GetDB().Create(&news)
//	return &news, res.Error
//}
//
//func (n *NewsScrapResult) Create() error {
//	//news := NewsScrapResult{Href: "https://naver.com", Title: "NewTitle"}
//	//news.Create()
//	res := db.Create(&n)
//	return res.Error
//}
//
//func (n *NewsScrapResult) Upsert() error {
//	res := db.Save(n)
//	return res.Error
//}
