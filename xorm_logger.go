package database

import (
	`fmt`

	`github.com/pangum/logging`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
	`xorm.io/xorm/log`
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

func (l *xormLogger) Debug(v ...interface{}) {
	if l.showed {
		l.logger.Debug(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
	}
}

func (l *xormLogger) Debugf(format string, v ...interface{}) {
	if l.showed {
		l.logger.Debug(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
	}
}

func (l *xormLogger) Info(v ...interface{}) {
	if l.showed {
		l.logger.Info(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
	}
}

func (l *xormLogger) Infof(format string, v ...interface{}) {
	if l.showed {
		l.logger.Info(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
	}
}

func (l *xormLogger) Warn(v ...interface{}) {
	if l.showed {
		l.logger.Warn(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
	}
}

func (l *xormLogger) Warnf(format string, v ...interface{}) {
	if l.showed {
		l.logger.Warn(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
	}
}

func (l *xormLogger) Error(v ...interface{}) {
	if l.showed {
		l.logger.Error(`数据库`, field.String(`xorm`, fmt.Sprint(v...)))
	}
}

func (l *xormLogger) Errorf(format string, v ...interface{}) {
	if l.showed {
		l.logger.Error(`数据库`, field.String(`xorm`, fmt.Sprintf(format, v...)))
	}
}

func (l *xormLogger) Level() (level log.LogLevel) {
	switch l.logger.Level() {
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

func (l *xormLogger) SetLevel(level log.LogLevel) {
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
	l.logger.Sets(simaqian.Levels(lvl))
}

func (l *xormLogger) ShowSQL(show ...bool) {
	if 0 == len(show) {
		l.showed = true
	} else {
		l.showed = show[0]
	}
}

func (l *xormLogger) IsShowSQL() bool {
	return l.showed
}
