package database

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestRepository_FindAll(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	for i := 1; i <= 10; i++ {
		if err := repo.Create(&NewsScrapResult{
			Href:  "www.host.com",
			Title: fmt.Sprintf("Title %d", i),
		}); err != nil {
			t.Error(err)
		}
	}

	var findNewsList []NewsScrapResult
	if err := repo.FindAll(&findNewsList); err != nil {
		t.Error(err)
	}

	for _, n := range findNewsList {
		fmt.Println(n)
	}
}

func TestRepository_FindById(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	news := NewsScrapResult{Href: "www.host.com", Title: fmt.Sprintf("Title")}
	if err := repo.Create(&news); err != nil {
		t.Error(err)
	}
	fmt.Println("Create:", news)

	findNews := NewsScrapResult{}
	err := repo.FindById(&findNews, news.Id)
	fmt.Println("Find:", findNews)

	assert.Equal(t, err, nil)
}

func TestRepository_FindPaged(t *testing.T) {
	suite := createTestSuite()
	repo := NewRepository(suite.db)

	findNewsList := []NewsScrapResult{}
	err := repo.FindPaged(&findNewsList, 3, 3)
	if err != nil {
		t.Error(err)
	}

	for _, n := range findNewsList {
		fmt.Println(n)
	}
}

func TestCreateMockRepository(t *testing.T) {
	repo := CreateMockRepository()
	fmt.Println(repo)

	news := NewsScrapResult{Href: "www.host.com", Title: fmt.Sprintf("TestCreateMockRepository")}
	err := repo.Create(&news)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Create:", news)
}
