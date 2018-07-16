package main

import (
	"time"
	"net/http"
	"io/ioutil"
	"fmt"
	"compress/gzip"
	"bytes"
)

//var url = "http://fooddrug.service-alpha.wochacha.cn/tasklist/test"
var url = "http://192.168.200.222:9000/register?barcode=6970872410006"
var url2 = "http://192.168.200.222:9000/pure_random?barcode=6970872410006&count=1"

func main() {
	for i := 0; i < 5; i++{
		go oper(i)
	}
	time.Sleep(1*time.Second)
}

func oper( arg interface{}) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	resp2, _ := http.Get(url2)
	defer resp2.Body.Close()
	body2, _ := ioutil.ReadAll(resp2.Body)


	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(body2)

	w.Flush()
	fmt.Println("gzip size:", len(b.Bytes()))

	r, _ := gzip.NewReader(&b)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	fmt.Println("ungzip size:", len(undatas), string(undatas))

}
