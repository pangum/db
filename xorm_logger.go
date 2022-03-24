package database

import (
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/pangum/logging"
	"github.com/storezhang/simaqian"
	"xorm.io/xorm/log"
)

type xormLogger struct {
	logger *logging.Logger
	showed bool
}

func newXormLogger(logger *logging.Logger) *xormLogger {
	return &xormLogger{
		logger: logger,
		showed: false,
	}
}

func (xl *xormLogger) Debug(v ...interface{}) {
	xl.logger.Debug(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
}

func (xl *xormLogger) Debugf(format string, v ...interface{}) {
	xl.logger.Debug(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
}

func (xl *xormLogger) Info(v ...interface{}) {
	xl.logger.Info(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
}

func (xl *xormLogger) Infof(format string, v ...interface{}) {
	xl.logger.Info(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
}

func (xl *xormLogger) Warn(v ...interface{}) {
	xl.logger.Warn(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
}

func (xl *xormLogger) Warnf(format string, v ...interface{}) {
	xl.logger.Warn(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
}

func (xl *xormLogger) Error(v ...interface{}) {
	xl.logger.Error(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
}

func (xl *xormLogger) Errorf(format string, v ...interface{}) {
	xl.logger.Error(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
}

func (xl *xormLogger) Level() (level log.LogLevel) {
	switch xl.logger.Level() {
	case simaqian.LevelDebug:
		level = log.LOG_DEBUG
	case simaqian.LevelInfo:
		level = log.LOG_INFO
	case simaqian.LevelWarn:
		level = log.LOG_WARNING
	case simaqian.LevelError:
		level = log.LOG_ERR
	default:
		level = log.LOG_UNKNOWN
	}

	return
}

func (xl *xormLogger) SetLevel(level log.LogLevel) {
	var lvl string
	switch level {
	case log.LOG_DEBUG:
		lvl = `debug`
	case log.LOG_INFO:
		lvl = `info`
	case log.LOG_WARNING:
		lvl = `warn`
	case log.LOG_ERR:
		lvl = `error`
	}
	xl.logger.Sets(simaqian.Levels(lvl))
}

func (xl *xormLogger) ShowSQL(show ...bool) {
	if 0 == len(show) {
		xl.showed = true
	} else {
		xl.showed = show[0]
	}
}

func (xl *xormLogger) IsShowSQL() bool {
	return xl.showed
}
