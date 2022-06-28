package middleware

import (
	"fmt"
	"order-food-service/global"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/**
1. 初始化导入的包（递归导入）
2. 对包块中声明的变量进行计算和分配初始值
3. 执行包中的init函数
4. init函数在main函数之前执行
*/
func init() {
	initAppLog()
	initWebLog()
}

/**
初始化AppLog
*/
func initAppLog() {
	logFileName := "app.log"
	global.AppLog = initLog(logFileName)
}

/**
初始化WebLog
*/
func initWebLog() {
	logFileName := "access.log"
	global.WebLog = initLog(logFileName)
}

/**
初始化日志句柄
*/
func initLog(logFileName string) *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logPath := "./"
	logName := logPath + logFileName
	var f *os.File
	var err error
	//判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		fmt.Println("open log file failed")
	}

	log.Out = f
	log.Level = logrus.InfoLevel
	return log
}

/**
Gin中间件函数，记录请求日志
*/
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		global.WebLog.WithFields(logrus.Fields{
			"http_status": statusCode,
			"total_time":  latencyTime,
			"ip":          clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}).Info("access")
	}
}
