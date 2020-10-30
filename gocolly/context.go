package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func main() {
	//uriExp,err := regexp.MustCompile(`http://go-colly.org/docs/.*?`)
	// Instantiate default collector
	c := colly.NewCollector(
		//colly.AllowedDomains("go-colly.org"),
		colly.URLFilters(regexp.MustCompile(`http://go-colly.org/docs/.*?`)),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		e.Request.Ctx.Put("urlTitle", e.Text)
		e.Request.Ctx.Put("url", e.Attr("href"))
	})

	// Before making a request put the URL with
	// the key of "url" into the context of the request
	c.OnRequest(func(r *colly.Request) {
		//url := r.URL.ToPredict()
		//r.Ctx.Put("url", r.URL.ToPredict())
	})

	// After making a request get "url" from
	// the context of the request
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Ctx.Get("urlTitle"), r.Ctx.Get("url"))
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit("http://go-colly.org/docs/")
}
