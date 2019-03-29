package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main() {
	fName := "jianshu.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"标题", "作者", "时间", "头像", "内容2"})

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("jianshu.com"),
		colly.MaxDepth(1),
		//colly.URLFilters(
		//	regexp.MustCompile(`https://www.jianshu.com/p/\w+^[#comments]`),
		//	regexp.MustCompile(`https://www.jianshu.com/u/.*?`),
		//	),
		)
	// Create another collector to scrape course details
	detailCollector := c.Clone()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		hrefByte := []byte(href)
		if len(hrefByte)>3 && href[:3]=="/p/" {
			fmt.Println(href)
			e.Request.Visit(href)
			//return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	detailCollector.OnHTML("div.article", func(e *colly.HTMLElement) {
		a := []string{
			e.ChildText("h1.title"),
			e.ChildText("div.author .name a"),
			e.ChildText("div.author .publish-time"),
			e.ChildAttr("div.author .avatar img", "src"),
			e.ChildText("div.show-content .show-content-free"),
		}
		writer.Write(a)
	})

	//detailCollector.OnRequest(func(r *colly.Request) {
	//	fmt.Println(r.URL)
	//})

	c.Visit("https://www.jianshu.com/u/e0be3b4f7800")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
