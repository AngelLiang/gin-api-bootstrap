package util

import (
	// "os"
	"time"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger

func GinLogrus(log *logrus.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 开始请求计时
        start := time.Now()

        // 处理请求
        c.Next()

        // 记录请求和响应信息
        log.WithFields(logrus.Fields{
            "status":  c.Writer.Status(),
            "method":  c.Request.Method,
            "path":    c.Request.URL.Path,
            "userAgent": c.Request.UserAgent(),
            "ip":      c.ClientIP(),
            "latency": time.Since(start),
        }).Info("Gin request")

        // 清除Gin的错误信息
        c.Errors = nil
    }
}

func InitLogrus() *logrus.Logger {
    log := logrus.New()
	intLevel := logrus.DebugLevel
	// switch level {
	// 	case "error":
	// 		intLevel = logrus.ErrorLevel
	// 	case "warning":
	// 		intLevel = logrus.WarningLevel
	// 	case "info":
	// 		intLevel = logrus.InformationalLevel
	// 	case "debug":
	// 		intLevel = logrus.DebugLevel
	// }

    // 设置日志级别
    log.SetLevel(intLevel)
    
    logFile := &lumberjack.Logger{
        Filename:   "log/log.log", // 日志文件路径
        MaxSize:    500,                      // 单个日志文件的最大尺寸（以MB为单位）
        MaxBackups: 10,                        // 保留的旧日志文件的最大个数
        MaxAge:     30,                       // 保留的旧日志文件的最大天数
        Compress:   true,                     // 是否压缩旧日志文件
    }

    // // 将日志记录到标准输出
    log.SetOutput(logFile)
    // log.SetOutput(os.Stdout)

	logger = log
    return logger
}

// Log 返回日志对象
func Log() *logrus.Logger {
	return logger
}
