package util

import (
	"os"
	"time"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
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
    
    // 将日志记录到标准输出
    log.SetOutput(os.Stdout)

	logger = log
    return logger
}

// Log 返回日志对象
func Log() *logrus.Logger {
	return logger
}
