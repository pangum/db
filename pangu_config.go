package db

type panguConfig struct {
	// 关系型数据库配置
	Db *config `json:"db" yaml:"db" xml:"db" toml:"db" validate:"required"`
}
