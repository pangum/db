package internal

import (
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/pangum/logging"
	"xorm.io/xorm/log"
)

type Xorm struct {
	logger logging.Logger
	showed bool
}

func NewXorm(logger logging.Logger) *Xorm {
	return &Xorm{
		logger: logger,
		showed: false,
	}
}

func (xl *Xorm) Debug(v ...interface{}) {
	xl.logger.Debug("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (xl *Xorm) Debugf(format string, v ...any) {
	xl.logger.Debug("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (xl *Xorm) Info(v ...any) {
	xl.logger.Info("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (xl *Xorm) Infof(format string, v ...any) {
	xl.logger.Info("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (xl *Xorm) Warn(v ...any) {
	xl.logger.Warn("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (xl *Xorm) Warnf(format string, v ...any) {
	xl.logger.Warn("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (xl *Xorm) Error(v ...any) {
	xl.logger.Error("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (xl *Xorm) Errorf(format string, v ...any) {
	xl.logger.Error("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (xl *Xorm) Level() (level log.LogLevel) {
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

func (xl *Xorm) SetLevel(level log.LogLevel) {
	var lvl string
	switch level {
	case log.LOG_DEBUG:
		lvl = "debug"
	case log.LOG_INFO:
		lvl = "info"
	case log.LOG_WARNING:
		lvl = "warn"
	case log.LOG_ERR:
		lvl = "error"
	}
	xl.logger.Enable(simaqian.ParseLevel(lvl))
}

func (xl *Xorm) ShowSQL(show ...bool) {
	if 0 == len(show) {
		xl.showed = true
	} else {
		xl.showed = show[0]
	}
}

func (xl *Xorm) IsShowSQL() bool {
	return xl.showed
}
