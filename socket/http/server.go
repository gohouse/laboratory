package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type persons struct {
	NAME string
	AGE  int
}

func main() {
	fmt.Println("server start ......")
	http.HandleFunc("/test", hand)
	err := http.ListenAndServe(":1026", nil)
	if err != nil {
		fmt.Println("err")
	}
}

func hand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
	w.Write([]byte("connect ok"))

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read err")
	}
	fmt.Println(body)
	fmt.Println(string(body))
	tmp := persons{}
	err = json.Unmarshal(body, &tmp) //JSON还原

	if err != nil {
		fmt.Println("unmarshal err")
	}
	fmt.Println(tmp)
}
