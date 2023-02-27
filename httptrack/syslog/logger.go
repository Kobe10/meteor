package syslog

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Logger interface {
	InitContext(ctx *gin.Context)
	withContext() *log.Entry
	Error(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Trace(...interface{})
	Debug(...interface{})
	Panic(...interface{})
	Fatal(...interface{})
	WithFields(fields map[string]interface{}) *log.Entry
	WithField(key string, value interface{}) *log.Entry
}

//拥有此类便拥有了写日志的能力
type requestLogger struct {
	requestId interface{}
}

//获取日志类 带有请求唯一标识
func NewLogger(ctx *gin.Context) Logger {
	var l Logger = new(requestLogger)
	l.InitContext(ctx)
	return l
}

//防止报错使用
func newEmptyLogger() Logger {
	return new(requestLogger)
}

//初始化上下文
func (l *requestLogger) InitContext(ctx *gin.Context) {
	if ctx != nil {
		if ctx.Request != nil {
			l.requestId = ctx.Request.Context().Value(RequestIdKey)
		}
		if l.requestId == nil {
			l.requestId = ctx.Value(RequestIdKey)
		}
	}
}

//携带上下文信息
func (l *requestLogger) withContext() *log.Entry {
	return GetLogger().WithField("id", l.requestId)
}

//带请求唯一标识的日志
func (l *requestLogger) Error(s ...interface{}) {
	l.withContext().Error(s...)
}

func (l *requestLogger) Warn(s ...interface{}) {
	l.withContext().Warn(s...)
}

func (l *requestLogger) Info(s ...interface{}) {
	l.withContext().Info(s...)
}

func (l *requestLogger) Trace(s ...interface{}) {
	l.withContext().Trace(s...)
}

func (l *requestLogger) Debug(s ...interface{}) {
	l.withContext().Debug(s...)
}

func (l *requestLogger) Panic(s ...interface{}) {
	l.withContext().Panic(s...)
}

func (l *requestLogger) Fatal(s ...interface{}) {
	l.withContext().Fatal(s...)
}

//输入自定义字段 批量
func (l *requestLogger) WithFields(fields map[string]interface{}) *log.Entry {
	return l.withContext().WithFields(fields)
}

//输入自定义字段 单个
func (l *requestLogger) WithField(key string, value interface{}) *log.Entry {
	return l.withContext().WithField(key, value)
}
