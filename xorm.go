package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/elliotchance/sshtunnel"
	"github.com/goexl/gox"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
	"golang.org/x/crypto/ssh"
	"xorm.io/core"
	"xorm.io/xorm"
)

// 创建Xorm操作引擎
func newEngine(config *pangu.Config, logger logging.Logger) (engine *Engine, err error) {
	wrap := new(wrapper)
	if err = config.Load(wrap); nil != err {
		return
	}
	db := wrap.Db

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

func newXorm(db *config, logger logging.Logger) (engine *Engine, err error) {
	if se := enableSSH(db, logger); nil != se {
		err = se
	} else if dsn, de := db.dsn(); nil != de {
		err = de
	} else {
		engine = new(Engine)
		engine.Engine, err = xorm.NewEngine(db.Type, dsn)
	}

	return
}

func enableSSH(db *config, logger logging.Logger) (err error) {
	if nil == db.SSH && !db.SSH.Enable() {
		return
	}

	password := db.SSH.Password
	keyfile := db.SSH.Keyfile
	auth := gox.Ift("" != password, ssh.Password(password), sshtunnel.PrivateKeyFile(keyfile))
	host := fmt.Sprintf("%s@%s", db.SSH.Username, db.SSH.Addr)
	if tunnel, ne := sshtunnel.NewSSHTunnel(host, auth, db.Addr, "65512"); nil != ne {
		err = ne
	} else {
		tunnel.Log = newSSHLogger(logger)
		go startTunnel(tunnel)
		time.Sleep(100 * time.Millisecond)
		db.Addr = fmt.Sprintf("127.0.0.1:%d", tunnel.Local.Port)
	}

	return
}

func startTunnel(tunnel *sshtunnel.SSHTunnel) {
	_ = tunnel.Start()
}
