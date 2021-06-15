package sql

import (
	`strings`

	`github.com/storezhang/pangu`
	`xorm.io/core`
	`xorm.io/xorm`
	`xorm.io/xorm/log`
)

// 创建Xorm操作引擎
func newXormEngine(config *pangu.Config) (engine *xorm.Engine, err error) {
	var panguConfig panguConfig
	if err = config.Struct(panguConfig); nil != err {
		return
	}
	sql := panguConfig.Sql

	var dsn string
	if dsn, err = sql.dsn(); nil != err {
		return
	}

	if engine, err = xorm.NewEngine(sql.Type, dsn); nil != err {
		return
	}

	// 调试模式下打开各种可调试的选项
	if sql.Show {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(log.LOG_DEBUG)
	}

	// 配置数据库连接池
	engine.SetMaxOpenConns(sql.Connection.MaxOpen)
	engine.SetMaxIdleConns(sql.Connection.MaxIdle)
	engine.SetConnMaxLifetime(sql.Connection.MaxLifetime)

	// 测试数据库连接成功
	if sql.Ping {
		if err = engine.Ping(); nil != err {
			return
		}
	}

	// 设置名称转换（列名及表名）
	core.NewCacheMapper(core.GonicMapper{})
	if "" != strings.TrimSpace(sql.Prefix) {
		core.NewPrefixMapper(core.GonicMapper{}, sql.Prefix)
	}
	if "" != strings.TrimSpace(sql.Suffix) {
		core.NewSuffixMapper(core.GonicMapper{}, sql.Suffix)
	}

	return
}
