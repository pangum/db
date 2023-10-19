package db

import (
	"github.com/goexl/gox"
	"xorm.io/xorm"
)

type Engine struct {
	*xorm.Engine
	gox.CannotCopy
}
