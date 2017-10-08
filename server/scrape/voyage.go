package scrape

import (
	"fmt"
	"strconv"

	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/RyomaK/tech_news/server/model"
)

func Voyage(year int) []model.Article {
	url := "http://techlog.voyagegroup.com/archive/" + strconv.Itoa(year)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("voyage scarapping failed")
	}
	articles := []model.Article{}
	doc.Find("section").Each(func(_ int, s *goquery.Selection) {
		article := model.Article{Company: "voyage"}
		article.Title = s.Find(".entry-title-link").Text()
		t, _ := time.Parse("2006-01-02", s.Find("time").Text())
		article.PostedAt = t
		article.URL, _ = s.Find(".entry-title-link").Attr("href")
		article.Text = s.Find(".entry-description").Text()
		articles = append(articles, article)
	})

	return articles
}

func NewVoyage(year int) model.Article {
	url := "http://techlog.voyagegroup.com/archive/" + strconv.Itoa(year)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("voyage scarapping failed")
	}
	article := model.Article{Company: "voyage"}
	doc.Find("section").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		article.Title = s.Find(".entry-title-link").Text()
		t, _ := time.Parse("2006-01-02", s.Find("time").Text())
		article.PostedAt = t
		article.URL, _ = s.Find(".entry-title-link").Attr("href")
		article.Text = s.Find(".entry-description").Text()
		return false
	})

	return article
}
