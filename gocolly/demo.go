package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("go-colly.org"),
		)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasPrefix(href, "/docs/") {
			//href!="/docs/" {
			//fmt.Println(e.Text, e.Attr("href"))
			e.Request.Visit(href)
		}
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		if e.Request.URL.String() != "http://go-colly.org/docs/" {
			fmt.Println(e.Text)
		}

		//fmt.Println(e.Text, e.Request.URL.String())
		//fmt.Println(e.Text)
	})

	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})

	c.Visit("http://go-colly.org/")
}