package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func main() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("static/log/logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.Error("错误啊")
}
// "static/log/logrus.log"
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
