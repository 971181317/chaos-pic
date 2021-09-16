package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

const LogPath = "log"

// LogModel 日志类型
type LogModel byte

const (
	DEBUG   LogModel = 0 //测试使用，控制台和文件均输出
	PRODUCT LogModel = 1 //只输出文件
)

// SetLogModel 设置日志的类型
func SetLogModel(m LogModel) error {
	//判断文件夹
	if stat, err := os.Stat(LogPath); os.IsNotExist(err) {
		os.Mkdir(LogPath, 0777)
		os.Chmod(LogPath, 0777)
	} else {
		if !stat.IsDir() {
			os.Mkdir(LogPath, 0777)
			os.Chmod(LogPath, 0777)
		}
	}

	//创建文件
	f, err := os.Create("log/chaos-pic " + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		return err
	}

	switch m {
	case DEBUG:
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	case PRODUCT:
		// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(f)
	}
	return nil
}
