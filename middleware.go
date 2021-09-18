package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type LogInfo struct {
	Path      string `json:"path"`
	Method    string `json:"method"`
	StartTime string `json:"start_time"`
	RunTime   string `json:"run_time"`
	Status    int    `json:"status"`
}

// InitReq 初始化请求中间件
func InitReq(c *gin.Context) {
	// 请求前
	t := time.Now()
	//header中生成logid
	logid := getLogid(&t)
	c.Header("x-log", logid)
	c.Set("logid", logid)

	c.Next()

	// 请求后
	loginfo, _ := json.Marshal(LogInfo{
		Path:      c.Request.URL.Path,
		Method:    c.Request.Method,
		StartTime: t.Format("2006-01-02 15:04:05"),
		RunTime:   fmt.Sprintf("%dms", time.Since(t).Milliseconds()),
		Status:    c.Writer.Status(),
	})
	logs.Info(c, string(loginfo))
}

func getLogid(t *time.Time) string {
	return fmt.Sprintf("%d%p", t.UnixNano(), t)
}
