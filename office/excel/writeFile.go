package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/gohouse/gorose/utils"
	"errors"
)

func main() {

	filePath := "static/tmp/MyXLSXFile.xlsx"

	datas := []map[string]interface{}{
		{"a":1,"b":2,"c":3},
		{"a":11,"b":12,"c":13},
		{"a":"搞事情","b":22,"c":"汉字"},
		{"a":21,"b":22,"c":23},
	}
	ExportExcel(filePath, datas)

}

func ExportExcel(filePath string, datas []map[string]interface{}) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	if len(datas)==0 {
		return errors.New("数据为空")
	}

	// 添加表头
	var keys []string
	row = sheet.AddRow()
	for key,_ := range datas[0] {
		keys = append(keys, key)
		cell = row.AddCell()
		cell.Value = key
	}

	// 循环写入内容
	for _,item := range datas{
		row = sheet.AddRow()
		for _,v := range keys {
			cell = row.AddCell()
			cell.Value = utils.ParseStr(item[v])

			err := file.Save(filePath)
			if err != nil {
				fmt.Printf(err.Error())
				return err
			}
		}
	}

	return nil
}