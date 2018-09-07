package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type LogData struct {
	key string
	val interface{}
}

type SelfLogFormatter struct {

}

func getLogData() LogData  {
	return LogData{
		"key",
		"value",
	}
}

func simpleLog()  {
	logrus.Info("a simple log msg")
}

func fieldsLog()  {
	res := getLogData()
	logrus.WithFields(logrus.Fields{
		"field-name": "field-value",
		"result": res,
	}).Info("A walrus appears")
}

func formatLog()  {

}

func httpLog()  {
	entry := logrus.WithFields(logrus.Fields{"request_id": 1, "user_ip": "0.0.0.0"})
	entry.Info("http info")
	entry.Error("http error")
}

func fileLog()  {
	file, err := os.OpenFile("logging/logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	//file, err := os.OpenFile("logging/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		logrus.Errorf("open log file fail. error:%s", err.Error())
	}
	logrus.Info("a log msg write to file & stdout")
}

func init()  {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main()  {
	simpleLog()
	fieldsLog()
	formatLog()
	fileLog()
	httpLog()
	
}
