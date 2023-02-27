package syslog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	maxLogAge       = 7 * 24 * time.Hour
	logRotationTime = 24 * time.Hour
	defaultLogFile  = "app"
	defaultLogPath  = "/Users/fuzhiqiang/logs"
	RequestIdKey    = "requestIdKey"
)

var logger *log.Logger

type LoggerConfig struct {
	LogPath string
	LogFile string
	Env     string
}

//不需要携带唯一请求id的logger 不建议直接使用
func GetLogger() *log.Logger {
	return logger
}

func Setup(config LoggerConfig) {

	logger = log.New()

	//正式环境关闭标准输出
	//if config.Env == env.ZEnvProd {
	//	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//	if err != nil {
	//		fmt.Println("Open Src File err", err)
	//	} else {
	//		logger.SetOutput(bufio.NewWriter(src))
	//	}
	//}
	//设置日志级别
	logger.SetLevel(log.TraceLevel)

	//设置日志格式
	logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if config.LogFile == "" {
		config.LogFile = defaultLogFile
	}

	if config.LogPath == "" {
		config.LogPath = defaultLogPath
	}

	fullFilePath := fmt.Sprintf("%s", config.LogPath) + "/" + fmt.Sprintf("%s", config.LogFile)

	//日志分级记录
	logWriterAccess, _ := rotatelogs.New(
		//日志路径
		fullFilePath+".access.%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fullFilePath+".access.log"),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(maxLogAge),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(logRotationTime),
	)
	// 设置 rotatelogs
	logWriterInfo, _ := rotatelogs.New(
		//日志路径
		fullFilePath+".info.%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fullFilePath+".info.log"),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(maxLogAge),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(logRotationTime),
	)

	logWriterWarning, _ := rotatelogs.New(
		//日志路径
		fullFilePath+".warning.%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fullFilePath+".warning.log"),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(maxLogAge),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(logRotationTime),
	)

	logWriterError, _ := rotatelogs.New(
		//日志路径
		fullFilePath+".error.%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fullFilePath+".error.log"),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(maxLogAge),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(logRotationTime),
	)

	writeMap := lfshook.WriterMap{
		log.TraceLevel: logWriterAccess,
		log.InfoLevel:  logWriterInfo,
		log.DebugLevel: logWriterInfo,
		log.WarnLevel:  logWriterWarning,
		log.ErrorLevel: logWriterError,
		log.PanicLevel: logWriterError,
		log.FatalLevel: logWriterError,
	}

	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

}
