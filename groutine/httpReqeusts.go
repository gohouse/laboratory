package main

import (
	"time"
	"net/http"
	"io/ioutil"
	"fmt"
)

//var url = "http://fooddrug.service-alpha.wochacha.cn/tasklist/test"

func main() {
	//var url = "http://localhost:8002/gettaskinfo?id="
	var url = "http://localhost:8002/admin/GetTaskList"
	//var url2 = "http://localhost:8002/gettaskinfo?id="
	for i := 0; i < 50; i++{
		go oper(fmt.Sprint(url), i)
	}
	time.Sleep(10*time.Second)
}

func oper( url ...interface{}) {
	resp, _ := http.Get(url[0].(string))
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
	fmt.Println(url[1])

	//resp2, _ := http.Get(url2)
	//defer resp2.Body.Close()
	//body2, _ := ioutil.ReadAll(resp2.Body)
	//
	//
	//var b bytes.Buffer
	//w := gzip.NewWriter(&b)
	//defer w.Close()
	//w.Write(body2)
	//
	//w.Flush()
	//fmt.Println("gzip size:", len(b.Bytes()))
	//
	//r, _ := gzip.NewReader(&b)
	//defer r.Close()
	//undatas, _ := ioutil.ReadAll(r)
	//fmt.Println("ungzip size:", len(undatas), string(undatas))

}
