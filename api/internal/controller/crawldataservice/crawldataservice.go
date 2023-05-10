package crawldataservice

import (
	"GoReactProject/internal/model/crawldatamodel"

	"github.com/gocolly/colly/v2"
)

func CollyCrawl(request crawldatamodel.CrawlDataRequest) crawldatamodel.CrawlDataResponse {
	var result crawldatamodel.CrawlDataResponse
	list := Crawl(request.Url)
	result.Url = list
	return result
}

func Crawl(url string) []string {
	c := colly.NewCollector()

	var listSrc []string
	c.OnHTML("img", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		listSrc = append(listSrc, src)
	})

	c.Visit(url)
	return listSrc
}
