package internal

import (
	"fmt"

	"github.com/goexl/gox/field"
	"github.com/pangum/logging"
)

type Ssh struct {
	logger logging.Logger
}

func NewSsh(logger logging.Logger) *Ssh {
	return &Ssh{
		logger: logger,
	}
}

func (sl *Ssh) Printf(format string, args ...any) {
	sl.logger.Info("连接隧道", field.New("ssh", fmt.Sprintf(format, args...)))
}
