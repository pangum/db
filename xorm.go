package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/elliotchance/sshtunnel"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
	"golang.org/x/crypto/ssh"
	"xorm.io/core"
	"xorm.io/xorm"
)

// 创建Xorm操作引擎
func newEngine(config *pangu.Config, logger *logging.Logger) (engine *Engine, err error) {
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}
	db := _panguConfig.Db

	// 创建引擎
	if engine, err = newXorm(db, logger); nil != err {
		return
	}

	// 替换成统一的日志框架
	engine.SetLogger(newXormLogger(logger))
	// 调试模式下打开各种可调试的选项
	if db.Show {
		engine.ShowSQL()
	}

	// 配置数据库连接池
	engine.SetMaxOpenConns(db.Connection.Open)
	engine.SetMaxIdleConns(db.Connection.Idle)
	engine.SetConnMaxLifetime(db.Connection.Lifetime)

	// 测试数据库连接成功
	if db.Ping {
		if err = engine.Ping(); nil != err {
			return
		}
	}

	// 设置名称转换（列名及表名）
	core.NewCacheMapper(core.GonicMapper{})
	if "" != strings.TrimSpace(db.Prefix) {
		core.NewPrefixMapper(core.GonicMapper{}, db.Prefix)
	}
	if "" != strings.TrimSpace(db.Suffix) {
		core.NewSuffixMapper(core.GonicMapper{}, db.Suffix)
	}

	return
}

func newXorm(database config, logger *logging.Logger) (engine *Engine, err error) {
	if nil != database.SSH && database.SSH.Enable() {
		var auth ssh.AuthMethod
		if "" != database.SSH.Password {
			auth = ssh.Password(database.SSH.Password)
		} else {
			auth = sshtunnel.PrivateKeyFile(database.SSH.Keyfile)
		}
		host := fmt.Sprintf("%s@%s", database.SSH.Username, database.SSH.Addr)
		tunnel := sshtunnel.NewSSHTunnel(host, auth, database.Addr, `65512`)
		tunnel.Log = newSSHLogger(logger)
		go func() {
			err = tunnel.Start()
		}()

		time.Sleep(100 * time.Millisecond)
		database.Addr = fmt.Sprintf("127.0.0.1:%d", tunnel.Local.Port)
	}

	var dsn string
	if dsn, err = database.dsn(); nil != err {
		return
	}

	engine = new(Engine)
	engine.Engine, err = xorm.NewEngine(database.Type, dsn)

	return
}
