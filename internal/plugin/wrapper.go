package plugin

type Wrapper struct {
	// 关系型数据库配置
	Db *Config `json:"db" yaml:"db" xml:"db" toml:"db" validate:"required"`
}
