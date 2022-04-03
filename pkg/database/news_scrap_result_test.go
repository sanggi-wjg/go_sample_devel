package database

import (
	"fmt"
	"testing"
)

func TestCreateNewsScrapResult(t *testing.T) {
	suite := createTestSuite()

	news := NewsScrapResult{Href: "https://222.com", Title: "NewTitleIS"}
	res := suite.db.Create(&news)

	fmt.Println(news.Id, news.Href, news.Title)
	fmt.Println(res)
}

func TestNewsScrapResultList(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	newsList := &[]NewsScrapResult{}
	err := repo.FindAll(newsList)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(newsList)
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
