package database

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCreateNewsScrapResult(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	// given
	news := &NewsScrapResult{Href: "https://123123.com", Title: "NewTitleIS123123"}

	// when
	err := repo.Create(news)
	fmt.Println(news.Id, news.Href, news.Title)

	// then
	assert.Equal(t, err, nil)
}

func TestNewsScrapResultList(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	// given
	newsList := &[]NewsScrapResult{}

	// when
	err := repo.FindAll(newsList)
	fmt.Println(newsList)

	// then
	assert.Equal(t, err, nil)
}

func TestSuccessCaseNewsScrapResultFindById(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	// given
	news := &NewsScrapResult{Href: "https://123123.com", Title: "NewTitleIS123123"}
	err := repo.Create(news)
	fmt.Println("create", news)

	// when
	findNews := &NewsScrapResult{}
	err = repo.FindById(findNews, 1)
	fmt.Println("find", findNews)

	// then
	assert.Equal(t, err, nil)
}

func TestFailCaseNewsScrapResultFindById(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	// given
	news := &NewsScrapResult{}

	// when
	err := repo.FindById(news, 99999)
	fmt.Println(news)

	// then
	assert.Equal(t, err, ErrNotFound)
}
