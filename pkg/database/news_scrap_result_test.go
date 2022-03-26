package database

import (
	"fmt"
	"testing"
)

func TestCreateNewsScrapResult(t *testing.T) {
	suite := createSQLMock()

	// given
	news := NewsScrapResult{Href: "https://123123.com"}

	// when
	//suite.mock.MatchExpectationsInOrder(false)
	//suite.mock.ExpectBegin()

	suite.db.Create(&news)

	//suite.mock.ExpectCommit()
	//if res.Error != nil {
	//	t.Error(res.Error)
	//}

	//suite.mock.ExpectRollback()

	// then
	err := suite.mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
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
