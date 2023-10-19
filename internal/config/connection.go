package config

import (
	"time"
)

type Connection struct {
	// 最大打开连接数
	Open int `default:"150" yaml:"open" json:"open" xml:"open" toml:"open"`
	// 最大休眠连接数
	Idle int `default:"30" yaml:"idle" json:"idle" xml:"idle" toml:"idle"`
	// 每个连接最大存活时间
	Lifetime time.Duration `default:"5s" yaml:"lifetime" json:"lifetime" xml:"lifetime" toml:"lifetime"`
}
