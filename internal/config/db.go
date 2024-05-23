package config

import (
	"fmt"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/pangum/db/internal/config/internal"
	"github.com/pangum/db/internal/internal/constant"
	"xorm.io/xorm/names"
)

type DB struct {
	// 数据库类型
	// nolint:lll
	Type string `default:"sqlite3" json:"type,omitempty" yaml:"type" xml:"type" toml:"type" validate:"required,oneof=mysql sqlite3 mssql oracle psql"`

	// 地址，填写服务器地址
	// nolint:lll
	Addr string `default:"127.0.0.1:3306" json:"addr,omitempty" yaml:"addr" xml:"addr" toml:"addr" validate:"required,hostname_port"`
	// 授权，用户名
	Username string `json:"username,omitempty" yaml:"username" xml:"username" toml:"username"`
	// 授权，密码
	Password string `json:"password,omitempty" yaml:"password" xml:"password" toml:"password"`
	// 连接协议
	// nolint: lll
	Protocol string `default:"tcp" json:"protocol,omitempty" yaml:"protocol" xml:"protocol" toml:"password" validate:"required,oneof=tcp udp"`

	// 连接池配置
	Connection Connection `json:"connection,omitempty" yaml:"connection" xml:"connection" toml:"connection"`

	// 表名规则
	// nolint: lll
	Mapper string `default:"gonic" json:"mapper,omitempty" yaml:"mapper" xml:"mapper" toml:"mapper" validate:"required,oneof=snake same gonic"`
	// 表名的前缀
	Suffix string `json:"suffix,omitempty" yaml:"suffix" xml:"suffix" toml:"suffix"`
	// 表名后缀
	Prefix string `json:"prefix,omitempty" yaml:"prefix" xml:"prefix" toml:"prefix"`
	// 连接的数据库名
	Schema string `json:"schema,omitempty" yaml:"schema" xml:"schema" toml:"schema" validate:"required"`
	// 路径
	// nolint:lll
	Path string `default:"data.db" json:"path,omitempty" yaml:"path" xml:"path" toml:"path" validate:"required_if=Type sqlite3"`

	// 额外参数
	// nolint: lll
	Parameters internal.Parameters `default:"{'parseTime': true, 'loc': 'Local'}" json:"parameters,omitempty" yaml:"parameters" xml:"parameters" toml:"parameters"`
	// 是否连接时测试数据库连接是否完好
	Ping bool `default:"true" json:"ping,omitempty" yaml:"ping" xml:"ping" toml:"ping"`
	// 是否显示执行语句
	Show bool `default:"false" json:"show,omitempty" yaml:"show" xml:"show" toml:"show"`

	// 代理连接
	SSH *Ssh `json:"ssh,omitempty" yaml:"ssh" xml:"ssh" toml:"ssh"`
	// 同步
	Sync Sync `json:"sync,omitempty" yaml:"sync" xml:"sync" toml:"sync"`
	// 参数配置
	Sqlite internal.Sqlite `json:"sqlite,omitempty" yaml:"sqlite" xml:"sqlite" toml:"sqlite"`
}

func (d *DB) TableMapper() (mapper names.Mapper) {
	switch d.Mapper {
	case constant.Gonic:
		mapper = new(names.GonicMapper)
	case constant.Snake:
		mapper = new(names.SnakeMapper)
	case constant.Same:
		mapper = new(names.SameMapper)
	default:
		mapper = new(names.GonicMapper)
	}

	return
}

func (d *DB) DSN() (dsn string, err error) {
	switch strings.ToLower(d.Type) {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@%s(%s)", d.Username, d.Password, d.Protocol, d.Addr)
		if "" != strings.TrimSpace(d.Schema) {
			dsn = fmt.Sprintf("%s/%s", dsn, strings.TrimSpace(d.Schema))
		}
	case "sqlite":
		dsn = d.Path
	case "sqlite3":
		dsn = d.Path
		if "" != d.Username && "" != d.Password {
			d.Parameters[d.Sqlite.Name] = ""
			d.Parameters[d.Sqlite.User] = d.Username
			d.Parameters[d.Sqlite.Password] = d.Password
			d.Parameters[d.Sqlite.Crypt] = "sha512"
		}
	default:
		err = exception.New().Message("不支持的数据库类型").Field(field.New("type", d.Type)).Build()
	}
	if nil != err {
		return
	}

	// 增加参数
	parameters := d.Parameters.String()
	if "" != parameters {
		dsn = fmt.Sprintf("%s?%s", dsn, parameters)
	}

	return
}

func (d *DB) SSHEnabled() bool {
	return nil != d.SSH && d.SSH.Enable()
}
