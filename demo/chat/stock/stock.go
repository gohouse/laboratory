package stock

import (
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gookit/color"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ReadStock(stock string) string {
	rList, err := GetSinaStock(stock)
	if err != nil {
		return err.Error()
	}

	res := ""
	for _, item := range rList {
		res = formatForTerminal(item.Name, item.Price, item.Percent, item.TickSize)
	}
	return res
}

type SinaStock struct {
	Name     string
	Price    float64
	Percent  float64
	TickSize string
}

func GetSinaStock(sname string) (list []*SinaStock, err error) {
	urlAdress := fmt.Sprintf("http://hq.sinajs.cn/list=%s", sname)
	req, err := http.Get(urlAdress)
	if err != nil {
		return
	}
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	enc := mahonia.NewDecoder("gbk")
	content := enc.ConvertString(string(b))
	list = make([]*SinaStock, 0)
	for _, item := range strings.Split(content, ";") {
		if strings.TrimSpace(item) == "" {
			continue
		}
		index := strings.Index(item, "\"")
		rs := strings.Split(item[index+1:len(item)-1], ",")
		if len(rs) < 3 {
			err = errors.New("获取股票错误")
			return
		}
		price, _ := strconv.ParseFloat(rs[3], 10)
		oldPrice, _ := strconv.ParseFloat(rs[2], 10)
		zdfv := price - oldPrice
		zdf := zdfv / oldPrice * 100
		sinaStock := &SinaStock{
			Name:     rs[0],
			Price:    price,
			Percent:  zdfv,
			TickSize: fmt.Sprintf("(%.2f%%)", zdf),
		}
		list = append(list, sinaStock)
	}
	return
}

func formatForTerminal(name string, price float64, delta float64, percentage string) string {
	var deltaFormatted string
	var percentageFormatted = percentage

	// use like func
	red := color.FgRed.Render
	green := color.FgGreen.Render
	if delta > 0 {
		deltaFormatted = red(fmt.Sprintf("+%.2f", delta))
		percentageFormatted = red(fmt.Sprintf("%s", percentage))

		//deltaFormatted = fmt.Sprintf("\x1b[31m+%.2f\x1b[0m", delta)
		//percentageFormatted = fmt.Sprintf("\x1b[31m%s\x1b[0m", percentage)
	} else if delta < 0 {
		deltaFormatted = green(fmt.Sprintf("%.2f", delta))
		percentageFormatted = green(fmt.Sprintf("%s", percentage))
		//deltaFormatted = fmt.Sprintf("\x1b[32m%.2f\x1b[0m", delta)
		//percentageFormatted = fmt.Sprintf("\x1b[32m%s\x1b[0m", percentage)
	} else {
		deltaFormatted = fmt.Sprintf("%.2f", delta)
	}
	return fmt.Sprintf("%s %.2f %s %s", green(name), price, deltaFormatted, percentageFormatted)
}
