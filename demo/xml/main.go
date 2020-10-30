package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Row struct {
	XMLName  xml.Name `xml:"xml"`
	Row []User `xml:"row"`
}
type User struct {
	XMLName  xml.Name `xml:"row"`
	Expect   string   `xml:"expect,attr"`   // 读取expect属性
	Code     string   `xml:"code,attr"`     // 读取expect属性
	Opencode string   `xml:"opencode,attr"` // 读取expect属性
	Opentime string   `xml:"opentime,attr"` // 读取opentime属性
}

func main() {

	//url := "https://test.lightengine.com.cn/lottery/period/selectPeriod?sign=b6ef23ec7e04c3b9b5304c8e380219b3&appId=a116"
	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	body := []byte(`<?xml version="1.0" encoding="utf-8"?>
<xml rows="13" remain="9999hrs">
  <row expect="20200804313" code="180204113800009" opencode="02,06,03,05,09" opentime="2020-08-04 15:36:33"/>
  <row expect="202008040939" code="180204113800010" opencode="10,09,06,07,02" opentime="2020-08-04 15:38:30"/>
  <row expect="202008040939" code="180204113800014" opencode="10,04,03,05,08,02,07,09,01,06" opentime="2020-08-04 15:38:30"/>
  <row expect="202008040939" code="180204113800015" opencode="0,4,9,7,9" opentime="2020-08-04 15:38:30"/>
  <row expect="202008040939" code="180204113800017" opencode="26,18,04,23,13,44,39" opentime="2020-08-04 15:38:30"/>
  <row expect="20200804313" code="180204113800019" opencode="04,01,13,20,10,44,41" opentime="2020-08-04 15:36:33"/>
  <row expect="20200804188" code="180204113800020" opencode="37,03,05,27,10,36,16" opentime="2020-08-04 15:37:30"/>
  <row expect="20200804313" code="180204113800030" opencode="9,0,3,7,5" opentime="2020-08-04 15:36:33"/>
  <row expect="20200804188" code="180204113800031" opencode="2,8,4,4,6" opentime="2020-08-04 15:37:30"/>
  <row expect="202008040939" code="180204113800032" opencode="3,6,6" opentime="2020-08-04 15:38:30"/>
  <row expect="202008040939" code="180204113800033" opencode="20,07,03,13,09,04,14,10" opentime="2020-08-04 15:38:30"/>
  <row expect="202008040939" code="180204113800034" opencode="18,04,01,10,02,11,12,19" opentime="2020-08-04 15:38:30"/>
  <row expect="202008040313" code="180204113800035" opencode="1,1,1" opentime="2020-08-04 15:36:33"/>
</xml>
`)
	var vs Row
	err := xml.Unmarshal(body, &vs)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	log.Printf("%#v",vs)
}
