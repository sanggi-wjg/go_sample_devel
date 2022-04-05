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
	//for _, res := range scarpNewsResultList {
	//	fmt.Println(res)
	//}

	saveResultList := make([]error, len(scarpNewsResultList))
	channel := make(chan error)

	// when
	suite := database.CreateSuite()
	repo := database.NewRepository(suite.GetDB())

	for i := 0; i < len(scarpNewsResultList); i++ {
		go testSave(repo, scarpNewsResultList[i], channel)
	}
	for i := 0; i < len(scarpNewsResultList); i++ {
		saveResultList[i] = <-channel
	}
	fmt.Println(saveResultList)
	// then

}

func testSave(repo *database.Repository, newsResult NewsResult, c chan error) {
	err := repo.Upsert(&database.NewsScrapResult{
		Href:  newsResult.Href,
		Title: newsResult.Title,
	})
	if err != nil {
		c <- err
	}
	c <- nil
}
