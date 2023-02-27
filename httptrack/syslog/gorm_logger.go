package syslog

import (
	"context"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

var Default = New(gormlogger.Config{
	SlowThreshold: 100 * time.Millisecond,
	LogLevel:      gormlogger.Warn,
	Colorful:      true,
})

func New(config gormlogger.Config) gormlogger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%v] [rows:%d] %s"
		traceWarnStr = "%s\n[%v] [rows:%d] %s"
		traceErrStr  = "%s %s\n[%v] [rows:%d] %s"
	)

	if config.Colorful {
		infoStr = gormlogger.Green + "%s\n" + gormlogger.Reset + gormlogger.Green + "[info] " + gormlogger.Reset
		warnStr = gormlogger.Blue + "%s\n" + gormlogger.Reset + gormlogger.Magenta + "[warn] " + gormlogger.Reset
		errStr = gormlogger.Magenta + "%s\n" + gormlogger.Reset + gormlogger.Red + "[error] " + gormlogger.Reset
		traceStr = gormlogger.Green + "%s\n" + gormlogger.Reset + gormlogger.Yellow + "[%.3fms] " + gormlogger.Blue + "[rows:%d]" + gormlogger.Reset + " %s"
		traceWarnStr = gormlogger.Green + "%s\n" + gormlogger.Reset + gormlogger.RedBold + "[%.3fms] " + gormlogger.Yellow + "[rows:%d]" + gormlogger.Magenta + " %s" + gormlogger.Reset
		traceErrStr = gormlogger.RedBold + "%s " + gormlogger.MagentaBold + "%s\n" + gormlogger.Reset + gormlogger.Yellow + "[%.3fms] " + gormlogger.Blue + "[rows:%d]" + gormlogger.Reset + " %s"
	}

	return &glogger{
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type glogger struct {
	gormlogger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *glogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l glogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		NewCommonLogger(ctx).Info(msg, data)
	}
}

// Warn print warn messages
func (l glogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		NewCommonLogger(ctx).Warn(msg, data)
	}
}

// Error print error messages
func (l glogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		NewCommonLogger(ctx).Error(msg, data)
	}
}

// Trace print sql message
func (l glogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= gormlogger.Error:
			sql, rows := fc()
			NewCommonLogger(ctx).WithFields(map[string]interface{}{
				"took": float64(elapsed.Nanoseconds()) / 1e6,
				"rows": rows,
				"sql":  sql,
			}).Info(utils.FileWithLineNum())
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
			sql, rows := fc()
			NewCommonLogger(ctx).WithFields(map[string]interface{}{
				"took": float64(elapsed.Nanoseconds()) / 1e6,
				"rows": rows,
				"sql":  sql,
			}).Info(utils.FileWithLineNum())
		case l.LogLevel >= gormlogger.Info:
			sql, rows := fc()
			NewCommonLogger(ctx).WithFields(map[string]interface{}{
				"took": float64(elapsed.Nanoseconds()) / 1e6,
				"rows": rows,
				"sql":  sql,
			}).Info(utils.FileWithLineNum())
		}
	}
}
