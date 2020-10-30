package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("jianshu.com"),
		//colly.MaxDepth(2),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		fmt.Println(e.Text)
		//if strings.HasPrefix(href, "/p/") {
			e.Request.Visit(href)
		//}
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		//if strings.HasPrefix(e.Request.URL.ToPredict(), "https://jianshu.com/p/") {
		//	fmt.Println(e.Text)
		//}

		//fmt.Println(e.Text, e.Request.URL.ToPredict())
		fmt.Println(e.Text)
	})

	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})

	c.Visit("https://www.jianshu.com/u/48e0a32026ae")
}