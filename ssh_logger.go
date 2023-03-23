package db

import (
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/pangum/logging"
)

type sshLogger struct {
	logger logging.Logger
}

func newSSHLogger(logger logging.Logger) *sshLogger {
	return &sshLogger{
		logger: logger,
	}
}

func (sl *sshLogger) Printf(format string, args ...any) {
	sl.logger.Info("连接隧道", field.New("ssh", fmt.Sprintf(format, args...)))
}
