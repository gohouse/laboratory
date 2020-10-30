package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractTgPhotoUrl(channelId, messageId string) (imgs []string) {
	// Request the HTML page.
	//var urlstr = "https://t.me/s/pasay_channel/9"
	var urlstr = fmt.Sprintf("https://t.me/s/%s/%s", channelId, messageId)
	res, err := http.Get(urlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".tgme_widget_message_wrap .tgme_widget_message").Find(".tgme_widget_message_bubble a.tgme_widget_message_photo_wrap").Each(func(i int, s *goquery.Selection) {
		if id, exists := s.Attr("href"); exists {
			if strings.HasSuffix(id, fmt.Sprintf("/%s",messageId)) {
				//widgetWrap.Eq(i).Find(".tgme_widget_message_bubble .tgme_widget_message_photo_wrap").Each(func(i int, s *goquery.Selection) {
					v, _ := s.Attr("style")
					var strRegex = `(ht|f)tp(s?)\:\/\/[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*(:(0-9)*)*(\/?)([a-zA-Z0-9\-\.\?\,\'\/\\\+&amp;%\$#_]*)?` //请求参数结尾- 英文或数字和[]内的各种字符
					exp := regexp.MustCompile(strRegex)
					res := exp.FindAllString(v, -1)

					for _, v2 := range res {
						if v2 != "" {
							imgs = append(imgs, strings.TrimSuffix(v2, "'"))
						}
					}
					return
				//})
			}
		}
	})
	return
}

func main() {
	imgs := ExtractTgPhotoUrl("xiuche_channel", "23")
	fmt.Println(imgs)
}
