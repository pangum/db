package plugin

import (
	"github.com/pangum/db/internal/config"
)

type Wrapper struct {
	// 关系型数据库配置
	Db *config.DB `json:"db" yaml:"db" xml:"db" toml:"db" validate:"required"`
}
