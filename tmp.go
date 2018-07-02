package main

import (
	"fmt"
	"encoding/xml"
	"bytes"
	"io"
	"golang.org/x/net/html/charset"
	"time"
	"github.com/levigross/grequests"
)

func main() {
	// 总期数
	var totalResult = 100
	// 今天的期数数字 20180604
	nTime := time.Now()
	//yesTime := nTime.AddDate(0,0,-1)
	today := nTime.Format("20060102")

	// 获取今天
	res := GetResoultCode(today)

	// 如果今天的期数不够, 就获取昨天的
	lenRes := len(res)
	if lenRes<totalResult {
		yesTime := nTime.AddDate(0,0,-1)
		yesDay := yesTime.Format("20060102")
		res2 := GetResoultCode(yesDay)

		// 获取昨天的后边结果, 补齐今天的数量
		for i := 0; i < totalResult-lenRes; i++ {
			res = append(res, res2[i])
		}
	}

	fmt.Println(len(res))
	fmt.Println((res))
}


func GetResoultCode(date string) []RowItem {
	// 头信息
	options := &grequests.RequestOptions{}

	// 开启一个session
	session := grequests.NewSession(options)
	// 网站首页
	urlIndex := "http://kaijiang.500.com/ssc.shtml"
	// 接口
	url := "http://kaijiang.500.com/static/public/ssc/xml/qihaoxml/"+date+".xml"
	session.Get(urlIndex, options)
	doc,_ := session.Get(url, options)

	return resultXmlDecode(doc.String())
}


type RowItem struct {
	Expect string `xml:"expect,attr"`
	Opencode string `xml:"opencode,attr"`
	Opentime string `xml:"opentime,attr"`
}
type Row struct{
	Rows []RowItem `xml:"row"`
}
func resultXmlDecode(xmls string) []RowItem {

	//xmls := `'<?xml version="1.0" encoding="gb2312" ?><xml><row  expect="20180604051" opencode="1,0,0,6,3" opentime="2018-06-04 14:31:00"/><row  expect="20180604050" opencode="8,9,7,1,8" opentime="2018-06-04 14:20:50"/><row  expect="20180604049" opencode="5,6,0,2,9" opentime="2018-06-04 14:11:00"/><row  expect="20180604048" opencode="1,9,6,8,7" opentime="2018-06-04 14:01:00"/><row  expect="20180604047" opencode="9,0,2,0,7" opentime="2018-06-04 13:51:00"/><row  expect="20180604046" opencode="1,7,8,4,2" opentime="2018-06-04 13:40:50"/><row  expect="20180604045" opencode="1,3,6,4,9" opentime="2018-06-04 13:30:50"/><row  expect="20180604044" opencode="7,3,9,1,7" opentime="2018-06-04 13:21:00"/><row  expect="20180604043" opencode="8,9,7,0,9" opentime="2018-06-04 13:10:50"/><row  expect="20180604042" opencode="0,0,3,3,3" opentime="2018-06-04 13:00:50"/><row  expect="20180604041" opencode="5,3,0,5,8" opentime="2018-06-04 12:51:00"/><row  expect="20180604040" opencode="7,8,4,0,2" opentime="2018-06-04 12:41:00"/><row  expect="20180604039" opencode="0,7,0,2,4" opentime="2018-06-04 12:30:50"/><row  expect="20180604038" opencode="3,2,4,2,6" opentime="2018-06-04 12:21:00"/><row  expect="20180604037" opencode="1,0,9,0,5" opentime="2018-06-04 12:11:00"/><row  expect="20180604036" opencode="3,1,0,3,5" opentime="2018-06-04 12:00:50"/><row  expect="20180604035" opencode="8,4,8,7,1" opentime="2018-06-04 11:51:00"/><row  expect="20180604034" opencode="0,6,8,0,7" opentime="2018-06-04 11:41:00"/><row  expect="20180604033" opencode="8,2,8,4,3" opentime="2018-06-04 11:30:50"/><row  expect="20180604032" opencode="2,1,8,2,2" opentime="2018-06-04 11:20:50"/><row  expect="20180604031" opencode="4,1,6,9,0" opentime="2018-06-04 11:10:50"/><row  expect="20180604030" opencode="4,9,4,8,1" opentime="2018-06-04 11:00:50"/><row  expect="20180604029" opencode="7,5,2,3,9" opentime="2018-06-04 10:51:00"/><row  expect="20180604028" opencode="3,8,8,0,6" opentime="2018-06-04 10:41:00"/><row  expect="20180604027" opencode="2,1,0,2,4" opentime="2018-06-04 10:30:50"/><row  expect="20180604026" opencode="1,6,8,9,6" opentime="2018-06-04 10:21:00"/><row  expect="20180604025" opencode="7,9,1,4,1" opentime="2018-06-04 10:11:00"/><row  expect="20180604024" opencode="6,6,4,0,3" opentime="2018-06-04 10:01:00"/><row  expect="20180604023" opencode="1,5,1,0,8" opentime="2018-06-04 01:56:00"/><row  expect="20180604022" opencode="1,0,9,0,6" opentime="2018-06-04 01:50:50"/><row  expect="20180604021" opencode="0,4,0,5,9" opentime="2018-06-04 01:46:00"/><row  expect="20180604020" opencode="1,1,2,4,5" opentime="2018-06-04 01:40:50"/><row  expect="20180604019" opencode="2,0,0,9,4" opentime="2018-06-04 01:36:00"/><row  expect="20180604018" opencode="6,5,5,3,3" opentime="2018-06-04 01:30:50"/><row  expect="20180604017" opencode="1,7,4,0,6" opentime="2018-06-04 01:26:00"/><row  expect="20180604016" opencode="4,1,1,3,9" opentime="2018-06-04 01:21:00"/><row  expect="20180604015" opencode="9,6,5,9,9" opentime="2018-06-04 01:15:50"/><row  expect="20180604014" opencode="0,3,3,3,5" opentime="2018-06-04 01:11:00"/><row  expect="20180604013" opencode="3,5,4,5,9" opentime="2018-06-04 01:05:50"/><row  expect="20180604012" opencode="6,1,8,4,7" opentime="2018-06-04 01:01:00"/><row  expect="20180604011" opencode="0,5,3,5,5" opentime="2018-06-04 00:55:50"/><row  expect="20180604010" opencode="0,6,9,8,9" opentime="2018-06-04 00:51:00"/><row  expect="20180604009" opencode="5,6,7,2,1" opentime="2018-06-04 00:46:00"/><row  expect="20180604008" opencode="9,6,3,8,7" opentime="2018-06-04 00:40:50"/><row  expect="20180604007" opencode="4,4,3,7,1" opentime="2018-06-04 00:36:00"/><row  expect="20180604006" opencode="5,9,8,4,5" opentime="2018-06-04 00:31:00"/><row  expect="20180604005" opencode="9,9,3,1,5" opentime="2018-06-04 00:26:00"/><row  expect="20180604004" opencode="4,5,1,8,6" opentime="2018-06-04 00:21:00"/><row  expect="20180604003" opencode="8,7,6,3,0" opentime="2018-06-04 00:16:00"/><row  expect="20180604002" opencode="7,6,7,2,4" opentime="2018-06-04 00:10:50"/><row  expect="20180604001" opencode="7,4,3,2,3" opentime="2018-06-04 00:06:00"/><row  expect="20180603120" opencode="4,1,1,0,7" opentime="2018-06-04 00:00:50"/></xml>'`

	decoder := xml.NewDecoder(bytes.NewReader([]byte(xmls)))
	decoder.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		return charset.NewReaderLabel(c, i)
	}
	result := &Row{}
	decoder.Decode(result)

	return result.Rows
}
