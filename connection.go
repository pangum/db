package database

import (
	`time`
)

type connection struct {
	// 最大打开连接数
	MaxOpen int `default:"150" yaml:"maxOpen" json:"maxOpen"`
	// 最大休眠连接数
	MaxIdle int `default:"30" yaml:"maxIdle" json:"maxIdle"`
	// 每个连接最大存活时间
	MaxLifetime time.Duration `default:"5s" yaml:"maxLifetime" json:"maxLifetime"`
}
