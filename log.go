package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type Log struct {
	logger *log.Logger
}

var logs Log

func InitLog(logPath string) {
	//判断文件夹
	if stat, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, 0777)
		os.Chmod(logPath, 0777)
	} else {
		if !stat.IsDir() {
			os.Mkdir(logPath, 0777)
			os.Chmod(logPath, 0777)
		}
	}

	//创建文件
	f, err := os.OpenFile(logPath+"/chaos-pic "+time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalf("log file err %s", err.Error())
	}

	logs.logger = log.New(f, "", log.Ldate|log.Ltime)
}

func (l Log) Info(c *gin.Context, format string) {
	logid, _ := c.Get("logid")
	file, method := getMethodAndFile()
	l.logger.Printf("%s [INFO]  method=%s %s %s", file, method, logid, format)
}

func (l Log) Infof(c *gin.Context, format string, args ...interface{}) {
	logid, _ := c.Get("logid")
	file, method := getMethodAndFile()
	l.logger.Printf("%s [INFO] method=%s %s %s", file, method, logid, fmt.Sprintf(format, args))
}

func (l Log) Error(c *gin.Context, format string) {
	logid, _ := c.Get("logid")
	file, method := getMethodAndFile()
	l.logger.Printf("%s [ERROR]  method=%s %s %s", file, method, logid, format)
}

func (l Log) Errorf(c *gin.Context, format string, args ...interface{}) {
	logid, _ := c.Get("logid")
	file, method := getMethodAndFile()
	l.logger.Printf("%s [ERROR] method=%s %s %s", file, method, logid, fmt.Sprintf(format, args))
}

func getMethodAndFile() (string, string) {
	pc, file, line, _ := runtime.Caller(2)
	fa := strings.Split(file, "/")

	method := runtime.FuncForPC(pc).Name()
	ma := strings.Split(method, "/")

	return fmt.Sprintf("%s:%d", fa[len(fa)-1], line), ma[len(ma)-1]
}
