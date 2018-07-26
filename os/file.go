package main

import (
	"os"
	"fmt"
	"path"
)

func main() {
	excelFileName := "/var/tmp/static/20180719152859-idfas-3304.csv"
	//os.Remove("./1.txt")

	//下载链接写入文件
	html := "/var/tmp/static/index.html"

	fp2,_ := os.Create(html)

	defer fp2.Close()
	fp2.WriteString(fmt.Sprintf(`<a href="./%s">点击下载(%s)</a>`, path.Base(excelFileName), path.Base(excelFileName)))

}
