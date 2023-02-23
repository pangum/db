package db

import (
	"xorm.io/xorm"
)

// Engine 数据库引擎简单封装
type Engine struct {
	*xorm.Engine
}
