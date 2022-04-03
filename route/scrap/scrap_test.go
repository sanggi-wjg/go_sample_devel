package scrap

import (
	"fmt"
	"go_sample_devel/pkg/database"
	"testing"
)

func TestScarpNaverNews(t *testing.T) {
	scrapList := GetScrapNaverNewsResult()
	for _, res := range scrapList {
		fmt.Println(res)
	}
}

func TestSaveScarpResult(t *testing.T) {
	// given
	scarpNewsResultList := GetScrapNaverNewsResult()
	for _, res := range scarpNewsResultList {
		fmt.Println(res)
	}

	saveResultList := make([]bool, len(scarpNewsResultList))
	channel := make(chan bool)

	// when
	for i := 0; i < len(scarpNewsResultList); i++ {
		go save(i, scarpNewsResultList[i], channel)
	}
	for i := 0; i < 10; i++ {
		saveResultList[i] = <-channel
	}
	fmt.Println(saveResultList)
}

func save(i int, newsResult NewsResult, c chan bool) {
	defer close(c)

	suite := database.CreateSuite()
	repo := database.NewRepository(suite.GetDB())

	fmt.Println(i)
	err := repo.Create(&database.NewsScrapResult{
		Href:  newsResult.Href,
		Title: newsResult.Title,
	})
	if err != nil {
		c <- false
	}
	c <- true
}
