package syslog

import (
	"context"
	"github.com/gin-gonic/gin"
)

//组合次结构便拥有了写日志的能力
type LogHandler struct {
	logger Logger
}

//设置logger 对象实例化后记得调用此方法设置通用参数
func (s *LogHandler) SetLogger(ctx *gin.Context) {
	s.logger = NewLogger(ctx)
}

//获取logger 防止报错
func (s *LogHandler) GetLogger() Logger {
	if s.logger != nil {
		return s.logger
	}
	s.logger = newEmptyLogger()
	return s.logger
}

type CommonLogHandler struct {
	logger ICommonLogger
}

//设置logger 对象实例化后记得调用此方法设置通用参数
func (s *CommonLogHandler) SetLogger(ctx context.Context) {
	s.logger = NewCommonLogger(ctx)
}

//获取logger 防止报错
func (s *CommonLogHandler) GetLogger() ICommonLogger {
	if s.logger != nil {
		return s.logger
	}
	s.logger = newEmptyCommonLogger()
	return s.logger
}
