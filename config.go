package database

import (
	`errors`
	`fmt`
	`strings`
)

// Config 数据库配置
type Config struct {
	// 数据库类型
	Type string `default:"sqlite3" json:"type" yaml:"type" validate:"required,oneof=mysql sqlite3 mssql oracle psql"`

	// 地址，填写服务器地址
	Addr string `default:"127.0.0.1:3306" json:"addr" validate:"required,hostname_port"`
	// 授权，用户名
	Username string `json:"username,omitempty" yaml:"username"`
	// 授权，密码
	Password string `json:"password,omitempty" yaml:"password"`
	// 连接协议
	Protocol string `default:"tcp" json:"protocol" yaml:"protocol" validate:"required,oneof=tcp udp"`

	// 连接池配置
	Connection ConnectionConfig `json:"connection" yaml:"connection"`

	// 表名的前缀
	Suffix string `json:"suffix,omitempty" yaml:"suffix"`
	// 表名后缀
	Prefix string `json:"prefix,omitempty" yaml:"prefix"`
	// 连接的数据库名
	Schema string `json:"schema" yaml:"schema" validate:"required"`

	// 额外参数
	Parameters string `json:"parameters,omitempty" yaml:"parameters"`
	// SQLite填写数据库文件的路径
	Path string `default:"data.db" json:"path,omitempty" yaml:"path"`
	// 是否连接时使用Ping测试数据库连接是否完好
	Ping bool `default:"true" json:"ping" yaml:"ping"`
	// 是否显示SQL执行语句
	Show bool `default:"false" json:"show" yaml:"show"`
}

func (c *Config) dsn() (dsn string, err error) {
	switch strings.ToLower(c.Type) {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@%s(%s)", c.Username, c.Password, c.Protocol, c.Addr)
		if "" != strings.TrimSpace(c.Schema) {
			dsn = fmt.Sprintf("%s/%s", dsn, strings.TrimSpace(c.Schema))
		}
	case "sqlite3":
		dsn = c.Path
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
