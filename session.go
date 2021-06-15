package database

import (
	`xorm.io/xorm`
)

// Session 描述一个事务
type Session struct {
	*xorm.Session
}
