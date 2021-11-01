package database

import (
	`strings`

	`github.com/pangum/pangu`
	`xorm.io/core`
	`xorm.io/xorm`
	`xorm.io/xorm/log`
)

// 创建Xorm操作引擎
func newXormEngine(config *pangu.Config) (engine *Engine, err error) {
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}
	database := _panguConfig.Database

	var dsn string
	if dsn, err = database.dsn(); nil != err {
		return
	}

	engine = new(Engine)
	if engine.Engine, err = xorm.NewEngine(database.Type, dsn); nil != err {
		return
	}

	// 调试模式下打开各种可调试的选项
	if database.Show {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(log.LOG_DEBUG)
	}

	// 配置数据库连接池
	engine.SetMaxOpenConns(database.Connection.MaxOpen)
	engine.SetMaxIdleConns(database.Connection.MaxIdle)
	engine.SetConnMaxLifetime(database.Connection.MaxLifetime)

	// 测试数据库连接成功
	if database.Ping {
		if err = engine.Ping(); nil != err {
			return
		}
	}

	// 设置名称转换（列名及表名）
	core.NewCacheMapper(core.GonicMapper{})
	if `` != strings.TrimSpace(database.Prefix) {
		core.NewPrefixMapper(core.GonicMapper{}, database.Prefix)
	}
	if `` != strings.TrimSpace(database.Suffix) {
		core.NewSuffixMapper(core.GonicMapper{}, database.Suffix)
	}

	return
}
