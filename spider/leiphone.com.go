package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("leiphone.com","www.leiphone.com"),
	)

	c.OnHTML(".lph-Nowsite a:nth-of-type(2)", func(e *colly.HTMLElement) {
		fmt.Printf("Link found: %s\n", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.leiphone.com/news/201702/zJaj9yNJuDzy4egF.html")
	//c.Visit("https://www.leiphone.com/")
}
