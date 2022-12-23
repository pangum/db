package database

import (
	"errors"
	"fmt"
	"strings"
)

type config struct {
	// 数据库类型
	// nolint:lll
	Type string `default:"sqlite3" json:"type" yaml:"type" xml:"type" toml:"type" validate:"required,oneof=mysql sqlite3 mssql oracle psql"`

	// 地址，填写服务器地址
	Addr string `default:"127.0.0.1:3306" json:"addr" yaml:"addr" xml:"addr" toml:"addr" validate:"required,hostname_port"`
	// 授权，用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
	// 授权，密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password"`
	// 连接协议
	Protocol string `default:"tcp" json:"protocol" yaml:"protocol" xml:"protocol" toml:"password" validate:"required,oneof=tcp udp"`

	// 连接池配置
	Connection connection `json:"connection" yaml:"connection" xml:"connection" toml:"connection"`

	// 表名的前缀
	Suffix string `json:"suffix" yaml:"suffix" xml:"suffix" toml:"suffix"`
	// 表名后缀
	Prefix string `json:"prefix" yaml:"prefix" xml:"prefix" toml:"prefix"`
	// 连接的数据库名
	Schema string `default:"data.db" json:"schema" yaml:"schema" xml:"schema" toml:"schema" validate:"required"`

	// 额外参数
	Parameters string `default:"parseTime=true" json:"parameters" yaml:"parameters" xml:"parameters" toml:"parameters"`
	// 是否连接时测试数据库连接是否完好
	Ping bool `default:"true" json:"ping" yaml:"ping" xml:"ping" toml:"ping"`
	// 是否显示执行语句
	Show bool `default:"false" json:"show" yaml:"show" xml:"show" toml:"show"`

	// SSH代理连接
	SSH *_ssh `json:"ssh" yaml:"ssh" xml:"ssh" toml:"ssh"`
}

func (c *config) dsn() (dsn string, err error) {
	switch strings.ToLower(c.Type) {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@%s(%s)", c.Username, c.Password, c.Protocol, c.Addr)
		if "" != strings.TrimSpace(c.Schema) {
			dsn = fmt.Sprintf("%s/%s", dsn, strings.TrimSpace(c.Schema))
		}
	case "sqlite3":
		dsn = c.Schema
	default:
		err = errors.New("不支持的数据库类型")
	}
	if nil != err {
		return
	}

	// 增加参数
	if "" != strings.TrimSpace(c.Parameters) {
		dsn = fmt.Sprintf("%s?%s", dsn, strings.TrimSpace(c.Parameters))
	}

	return
}
