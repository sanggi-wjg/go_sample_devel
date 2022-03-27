package database

import (
	"fmt"
	"testing"
)

func TestCreateNewsScrapResult(t *testing.T) {
	//suite := createTestSuite()

	news := NewsScrapResult{Href: "https://222.com", Title: "NewTitleIS"}
	//res := suite.db.Create(&news)
	err := news.Create()
	fmt.Println(news.Id, news.Href, news.Title)

	if err != nil {
		t.Error(err)
	}

	//db, mock, err := sqlmock.New()
	//if err != nil {
	//	t.Errorf("sql mock error: %v", err)
	//}
	//defer db.Close()
	//
	//news, err := CreateNewsScrapResult("https://123.com", "ThisIsTitle")
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(news)
	//
	//if err = mock.ExpectationsWereMet(); err != nil {
	//	t.Errorf("there were unfulfilled expectations: %s", err)
	//}
}

func TestGoCreateNewsScrapResults(t *testing.T) {
	//channelResults := make([]NewsScrapResult, 10)
	//channel := make(chan NewsScrapResult)
	//
	//for i := 0; i < 10; i++ {
	//	go test(i, channel)
	//}
	//for i := 0; i < 10; i++ {
	//	channelResults[i] = <-channel
	//}
	//fmt.Println(channelResults)
}

func test(i int, c chan NewsScrapResult) {
	defer close(c)

	fmt.Println(i)
	//news := NewsScrapResult{Href: "https://naver.com", Title: "NewTitle"}
	//news.Create()
	//c <- news
}
