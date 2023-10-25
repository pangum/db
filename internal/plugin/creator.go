package plugin

import (
	"fmt"
	"strings"
	"time"

	"github.com/elliotchance/sshtunnel"
	"github.com/goexl/gox"
	"github.com/goexl/log"
	"github.com/pangum/db/internal/db"
	"github.com/pangum/db/internal/internal"
	"github.com/pangum/pangu"
	"golang.org/x/crypto/ssh"
	"xorm.io/core"
	"xorm.io/xorm"
)

type Creator struct {
	// 用于解决命名空间
}

func (c *Creator) New(config *pangu.Config, logger log.Logger) (engine *db.Engine, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else if created, ne := c.new(wrapper.Db, logger); nil != ne {
		err = ne
	} else if se := c.setup(wrapper.Db, created, logger); nil != se {
		err = se
	} else {
		engine = created
	}

	return
}

func (c *Creator) NewTransaction(engine *db.Engine, logger log.Logger) *db.Transaction {
	return db.NewTransaction(engine, logger)
}

func (c *Creator) setup(config *Config, engine *db.Engine, logger log.Logger) (err error) {
	// 替换成统一的日志框架
	engine.SetLogger(internal.NewXorm(logger))
	// 调试模式下打开各种可调试的选项
	if config.Show {
		engine.ShowSQL()
	}

	// 配置数据库连接池
	engine.SetMaxOpenConns(config.Connection.Open)
	engine.SetMaxIdleConns(config.Connection.Idle)
	engine.SetConnMaxLifetime(config.Connection.Lifetime)

	// 设置名称转换（列名及表名）
	core.NewCacheMapper(core.GonicMapper{})
	if "" != strings.TrimSpace(config.Prefix) {
		core.NewPrefixMapper(core.GonicMapper{}, config.Prefix)
	}
	if "" != strings.TrimSpace(config.Suffix) {
		core.NewSuffixMapper(core.GonicMapper{}, config.Suffix)
	}

	// 测试数据库连接成功
	if config.Ping {
		err = engine.Ping()
	}

	return
}

func (c *Creator) new(config *Config, logger log.Logger) (engine *db.Engine, err error) {
	if se := c.enableSSH(config, logger); nil != se {
		err = se
	} else if dsn, de := config.dsn(); nil != de {
		err = de
	} else {
		engine = new(db.Engine)
		engine.Engine, err = xorm.NewEngine(config.Type, dsn)
	}

	return
}

func (c *Creator) enableSSH(conf *Config, logger log.Logger) (err error) {
	if !conf.sshEnable() {
		return
	}

	password := conf.SSH.Password
	keyfile := conf.SSH.Keyfile
	auth := gox.Ift("" != password, ssh.Password(password), sshtunnel.PrivateKeyFile(keyfile))
	host := fmt.Sprintf("%s@%s", conf.SSH.Username, conf.SSH.Addr)
	if tunnel, ne := sshtunnel.NewSSHTunnel(host, auth, conf.Addr, "65512"); nil != ne {
		err = ne
	} else {
		tunnel.Log = internal.NewSsh(logger)
		go c.startTunnel(tunnel)
		time.Sleep(100 * time.Millisecond)
		conf.Addr = fmt.Sprintf("127.0.0.1:%d", tunnel.Local.Port)
	}

	return
}

func (c *Creator) startTunnel(tunnel *sshtunnel.SSHTunnel) {
	_ = tunnel.Start()
}
