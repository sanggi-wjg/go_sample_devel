package scrap

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"go_sample_devel/pkg/database"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

type NewsResult struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type SaveNewsResult struct {
	NewsResult NewsResult `json:"newsResult"`
	SaveResult bool       `json:"saveResult"`
}

func (s *SaveNewsResult) SetResult(newsResult NewsResult, saveResult bool) {
	s.NewsResult = newsResult
	s.SaveResult = saveResult
}

func getNaverNewsDoc() (*goquery.Document, error) {
	res, err := http.Get("https://news.naver.com/main/main.naver?mode=LSD&mid=shm&sid1=100")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("status code is not 200:" + res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetScrapNaverNewsResult() []NewsResult {
	doc, err := getNaverNewsDoc()
	if err != nil {
		log.Fatalln(err)
	}
	var scrapList []NewsResult

	doc.Find("div.cluster_group").Find("li.cluster_item").Each(func(i int, items *goquery.Selection) {
		news := items.Find("a")
		href, _ := news.Attr("href")
		title := strings.TrimSpace(news.Text())
		if !utf8.ValidString(title) {
			title = ConvertToUTF8(title, "euc-kr")
		}

		scrapList = append(scrapList, NewsResult{href, title})
	})

	return scrapList
}

func GetNaverNews(c *gin.Context) {
	scrapList := GetScrapNaverNewsResult()

	c.JSON(200, gin.H{
		"message": "success",
		"result":  scrapList,
	})
}

func SaveNaverNews(c *gin.Context) {
	scrapList := GetScrapNaverNewsResult()
	scrapListSize := len(scrapList)

	saveNewsResult := make([]SaveNewsResult, scrapListSize)
	channel := make(chan bool)

	repo := database.NewRepository(database.GetDB())
	for i := 0; i < scrapListSize; i++ {
		go saveNews(repo, scrapList[i], channel)
	}
	for i := 0; i < scrapListSize; i++ {
		saveNewsResult[i].SetResult(scrapList[i], <-channel)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"result":  saveNewsResult,
	})
}

func saveNews(repo *database.Repository, newsResult NewsResult, c chan bool) {
	err := repo.Upsert(&database.NewsScrapResult{
		Href:  newsResult.Href,
		Title: newsResult.Title,
	})
	if err != nil {
		c <- false
	}
	c <- true
}
