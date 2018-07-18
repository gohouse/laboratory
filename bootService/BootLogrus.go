package bootService

import (
	"github.com/sirupsen/logrus"
	"os"
)
// NewLogrus 驱动动 logrus
// @param file 是否文件驱动,为空则默认控制台输出,为文件则输出到文件
// 		文件示例file[0] = "static/log/logrus.log"
func NewLogrus(file ...interface{}) *logrus.Logger {
	var log = logrus.New()
	log.Out = os.Stdout
	if len(file)>0 {
		if fileStr,ok := file[0].(string); ok && fileStr!=""{
			file, err := os.OpenFile(fileStr, os.O_CREATE|os.O_WRONLY, 0666)
			if err == nil {
				log.Out = file
			} else {
				log.Info("Failed to log to file, using default stderr")
			}
		}
	}
	return log
}
