package db

import (
	"xorm.io/xorm"
)

// Session 事务
type Session struct {
	*xorm.Session
}
