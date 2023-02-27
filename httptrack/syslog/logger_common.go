package syslog

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type ICommonLogger interface {
	InitContext(ctx context.Context)
	withContext() *log.Entry
	Error(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Debug(...interface{})
	Panic(...interface{})
	Fatal(...interface{})
	WithFields(fields map[string]interface{}) *log.Entry
	WithField(key string, value interface{}) *log.Entry
}

//拥有此类便拥有了写日志的能力
type commonLogger struct {
	ctxID interface{}
}

//获取日志类 带有请求唯一标识
func NewCommonLogger(ctx context.Context) ICommonLogger {
	var l ICommonLogger = new(commonLogger)
	l.InitContext(ctx)
	return l
}

//防止报错使用
func newEmptyCommonLogger() ICommonLogger {
	return new(commonLogger)
}

//初始化上下文
func (l *commonLogger) InitContext(ctx context.Context) {
	l.ctxID = ctx.Value(RequestIdKey)
}

//携带上下文信息
func (l *commonLogger) withContext() *log.Entry {
	return GetLogger().WithField("id", l.ctxID)
}

//带请求唯一标识的日志
func (l *commonLogger) Error(s ...interface{}) {
	l.withContext().Error(s...)
}

func (l *commonLogger) Warn(s ...interface{}) {
	l.withContext().Warn(s...)
}

func (l *commonLogger) Info(s ...interface{}) {
	l.withContext().Info(s...)
}

func (l *commonLogger) Debug(s ...interface{}) {
	l.withContext().Debug(s...)
}

func (l *commonLogger) Panic(s ...interface{}) {
	l.withContext().Panic(s...)
}

func (l *commonLogger) Fatal(s ...interface{}) {
	l.withContext().Fatal(s...)
}

//输入自定义字段 批量
func (l *commonLogger) WithFields(fields map[string]interface{}) *log.Entry {
	return l.withContext().WithFields(fields)
}

//输入自定义字段 单个
func (l *commonLogger) WithField(key string, value interface{}) *log.Entry {
	return l.withContext().WithField(key, value)
}
