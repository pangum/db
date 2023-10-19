package db

import (
	"github.com/goexl/gox"
	"xorm.io/xorm"
)

// Session 事务
type Session struct {
	*xorm.Session
	gox.CannotCopy
}
