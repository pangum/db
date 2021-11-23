package database

import (
	`fmt`
	log1 `log`
	`os`
	`strings`
	`time`

	`github.com/elliotchance/sshtunnel`
	`github.com/pangum/logging`
	`github.com/pangum/pangu`
	`golang.org/x/crypto/ssh`
	`xorm.io/core`
	`xorm.io/xorm`
	`xorm.io/xorm/log`
)

// 创建Xorm操作引擎
func newEngine(config *pangu.Config, logger *logging.Logger) (engine *Engine, err error) {
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}
	database := _panguConfig.Database

	// 创建引擎
	if engine, err = newXorm(database); nil != err {
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

func newXorm(database config) (engine *Engine, err error) {
	if nil != database.SSH && database.SSH.Enabled {
		var auth ssh.AuthMethod
		if `` != database.SSH.Password {
			auth = ssh.Password(database.SSH.Password)
		} else {
			auth = sshtunnel.PrivateKeyFile(database.SSH.Keyfile)
		}
		host := fmt.Sprintf(`%s@%s`, database.SSH.Username, database.SSH.Addr)
		tunnel := sshtunnel.NewSSHTunnel(host, auth, database.Addr, `65512`)
		tunnel.Log = log1.New(os.Stdout, "", log1.Ldate|log1.Lmicroseconds)
		go func() {
			err = tunnel.Start()
		}()

		time.Sleep(100 * time.Millisecond)
		database.Addr = fmt.Sprintf(`127.0.0.1:%d`, tunnel.Local.Port)
	}

	var dsn string
	if dsn, err = database.dsn(); nil != err {
		return
	}

	engine = new(Engine)
	engine.Engine, err = xorm.NewEngine(database.Type, dsn)

	return
}
