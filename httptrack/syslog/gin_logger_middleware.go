package syslog

import (
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		//上下文注入请求唯一标识设置请求ID
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), RequestIdKey, uuid.NewV4()))

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime).Milliseconds()

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式(2种不同方式)
		/*logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)*/
		NewLogger(c).WithFields(map[string]interface{}{
			"header":   c.Request.Header,
			"status":   statusCode,
			"took":     latencyTime,
			"clientIp": clientIP,
			"method":   reqMethod,
			"url":      reqUri,
		}).Trace()
	}
}
