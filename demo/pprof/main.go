package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	return sData
}

func main() {
	if true {
		go func() {
			for {
				log.Println(Add("https://github.com/EDDYCJY"))
			}
		}()
		runtime.GOMAXPROCS(1)               // 限制 CPU 使用数，避免过载
		runtime.SetMutexProfileFraction(1)  // 开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1)      // 开启对阻塞操作的跟踪
		http.ListenAndServe(":6060", nil)
	}

}