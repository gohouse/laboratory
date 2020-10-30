package main

import (
	"io/ioutil"
	"log"
	"net/http"
	u "net/url"
	"os"
)

func main() {
	var proxy *http.Response
	var err error
	if len(os.Args) > 1 {
		proxy, err = HttpGetFromProxy("http://172.24.0.2:8080/getip", "http://172.24.0.2:18080")
	} else {
		proxy, err = http.Get("http://172.24.0.2:8080/getip")
	}
	if err!=nil {
		log.Print("err:",err.Error())
		return
	}
	defer proxy.Body.Close()
	all, _ := ioutil.ReadAll(proxy.Body)
	log.Println("res: ", string(all))
}

// http GET 代理
func HttpGetFromProxy(url, proxyURL string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	proxy, err := u.Parse(proxyURL)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy : http.ProxyURL(proxy),
		},
	}
	return client.Do(req)
}
