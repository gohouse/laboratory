package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	main2()
}
func main2() {
	// memberid=&member_name=3Mw%2FeKBZ4Xt8Xz6f8EglMA%3D%3D&amount=qfTIDaJ/XlQdzF6QW/lBlA==&orderno=NBZZ1603635969717&bank=FXicugVowcRqv4lkYD1eMQ==&bank_name=&bank_number=UEpUQIwCk5MuX8EO5S97fQPAbY2MB5bBrYpUXLIlrMM=&bank_province=SqkOM6RawDOUGB49%2BvQeow%3D%3D&bank_city=SqkOM6RawDOUGB49%2BvQeow%3D%3D&payout_id=10262&payout_url=&user_ip=&user_name=uTvOuatRNysKoQjtWd+5jA==&number=&phone=f3D****Sg==&grade=&is_encryption=1&gosing=d008a172f3030b922065897b9450a801
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		log.Println(r.Form)
		fmt.Fprintln(w, "hello")
	})
	log.Println("http://127.0.0.1")
	http.ListenAndServe(":80", nil)
}

func fileServer()  {
	//文件浏览
	http.Handle("/", http.FileServer(http.Dir("../")))
	log.Println("http://localhost")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
