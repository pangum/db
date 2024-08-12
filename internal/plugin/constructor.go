package plugin

import (
	"fmt"
	"strings"
	"time"

	"github.com/elliotchance/sshtunnel"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/pangum/db/internal/config"
	"github.com/pangum/db/internal/db"
	"github.com/pangum/db/internal/internal"
	"github.com/pangum/pangu"
	"golang.org/x/crypto/ssh"
	"xorm.io/core"
	"xorm.io/xorm"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(config *config.DB, logger log.Logger) (engine *db.Engine, err error) {
	if created, ne := c.new(config, logger); nil != ne {
		err = ne
	} else if se := c.setup(config, created, logger); nil != se {
		err = se
	} else {
		engine = created
	}

	return
}

func (c *Constructor) NewTransaction(engine *db.Engine, logger log.Logger) *db.Transaction {
	return db.NewTransaction(engine, logger)
}

func (c *Constructor) NewSynchronizer(engine *db.Engine, config *config.DB, logger log.Logger) *db.Synchronizer {
	return db.NewSynchronizer(engine, config, logger)
}

func (c *Constructor) DB(config *pangu.Config) (db *config.DB, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		db = wrapper.Db
	}

	return
}

func (c *Constructor) setup(config *config.DB, engine *db.Engine, logger log.Logger) (err error) {
	// 替换成统一的日志框架
	engine.SetLogger(internal.NewXorm(logger))
	// 调试模式下打开各种可调试的选项
	if config.Verbose {
		engine.ShowSQL()
	}

	// 配置数据库连接池
	engine.SetMaxOpenConns(config.Connection.Open)
	engine.SetMaxIdleConns(config.Connection.Idle)
	engine.SetConnMaxLifetime(config.Connection.Lifetime)

	// 设置名称转换（列名及表名）
	mapper := config.TableMapper()
	core.NewCacheMapper(core.GonicMapper{})
	if "" != strings.TrimSpace(config.Prefix) {
		core.NewPrefixMapper(mapper, config.Prefix)
	}
	if "" != strings.TrimSpace(config.Suffix) {
		core.NewSuffixMapper(mapper, config.Suffix)
	}

	// 测试数据库连接成功
	if *config.Ping {
		logger.Info("开始测试数据库连接", field.New("config", config))
		err = engine.Ping()
	}

	return
}

func (c *Constructor) new(config *config.DB, logger log.Logger) (engine *db.Engine, err error) {
	if se := c.enableSSH(config, logger); nil != se {
		err = se
	} else if dsn, de := config.DSN(); nil != de {
		err = de
	} else {
		engine = new(db.Engine)
		engine.Engine, err = xorm.NewEngine(config.Type, dsn)
	}

	return
}

func (c *Constructor) enableSSH(conf *config.DB, logger log.Logger) (err error) {
	if !conf.SSHEnabled() {
		return
	}

	password := conf.SSH.Password
	keyfile := conf.SSH.Keyfile
	auth := gox.Ift("" != password, ssh.Password(password), sshtunnel.PrivateKeyFile(keyfile))
	host := fmt.Sprintf("%s@%s", conf.SSH.Username, conf.SSH.Addr)
	if tunnel, ne := sshtunnel.NewSSHTunnel(host, auth, conf.Host, "65512"); nil != ne {
		err = ne
	} else {
		tunnel.Log = internal.NewSsh(logger)
		go c.startTunnel(tunnel)
		time.Sleep(100 * time.Millisecond)
		conf.Host = fmt.Sprintf("127.0.0.1:%d", tunnel.Local.Port)
	}

	return
}

func (c *Constructor) startTunnel(tunnel *sshtunnel.SSHTunnel) {
	_ = tunnel.Start()
}
