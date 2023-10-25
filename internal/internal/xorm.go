package internal

import (
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	xorm "xorm.io/xorm/log"
)

type Xorm struct {
	logger log.Logger
	showed bool
}

func NewXorm(logger log.Logger) *Xorm {
	return &Xorm{
		logger: logger,
		showed: false,
	}
}

func (x *Xorm) Debug(v ...interface{}) {
	x.logger.Debug("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (x *Xorm) Debugf(format string, v ...any) {
	x.logger.Debug("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (x *Xorm) Info(v ...any) {
	x.logger.Info("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (x *Xorm) Infof(format string, v ...any) {
	x.logger.Info("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (x *Xorm) Warn(v ...any) {
	x.logger.Warn("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (x *Xorm) Warnf(format string, v ...any) {
	x.logger.Warn("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (x *Xorm) Error(v ...any) {
	x.logger.Error("数据库", field.New("xorm", fmt.Sprint(v...)))
}

func (x *Xorm) Errorf(format string, v ...any) {
	x.logger.Error("数据库", field.New("xorm", fmt.Sprintf(format, v...)))
}

func (x *Xorm) Level() (level xorm.LogLevel) {
	switch x.logger.Level() {
	case log.LevelDebug:
		level = xorm.LOG_DEBUG
	case log.LevelInfo:
		level = xorm.LOG_INFO
	case log.LevelWarn:
		level = xorm.LOG_WARNING
	case log.LevelError:
		level = xorm.LOG_ERR
	default:
		level = xorm.LOG_UNKNOWN
	}

	return
}

func (x *Xorm) SetLevel(level xorm.LogLevel) {
	var lvl string
	switch level {
	case xorm.LOG_DEBUG:
		lvl = "debug"
	case xorm.LOG_INFO:
		lvl = "info"
	case xorm.LOG_WARNING:
		lvl = "warn"
	case xorm.LOG_ERR:
		lvl = "error"
	}
	x.logger.Enable(log.ParseLevel(lvl))
}

func (x *Xorm) ShowSQL(show ...bool) {
	if 0 == len(show) {
		x.showed = true
	} else {
		x.showed = show[0]
	}
}

func (x *Xorm) IsShowSQL() bool {
	return x.showed
}
