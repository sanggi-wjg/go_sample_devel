package scrap

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

type NewsResult struct {
	Href, Title string
}

func NewNewsResult(href, title string) *NewsResult {
	s := NewsResult{Href: href, Title: title}
	return &s
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

func GetScrapNaverNewsResult() {
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

	for _, res := range scrapList {
		fmt.Println(res)
	}
}
