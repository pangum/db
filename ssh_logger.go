package database

import (
	`fmt`

	`github.com/pangum/logging`
	`github.com/storezhang/gox/field`
)

type sshLogger struct {
	logger *logging.Logger
}

func newSSHLogger(logger *logging.Logger) *sshLogger {
	return &sshLogger{
		logger: logger,
	}
}

func (sl *sshLogger) Printf(format string, v ...interface{}) {
	sl.logger.Info(`连接隧道`, field.String(`ssh`, fmt.Sprintf(format, v...)))
}
