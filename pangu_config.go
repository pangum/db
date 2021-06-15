package sql

type panguConfig struct {
	// 关系型数据库配置
	Sql Config `json:"sql" yaml:"sql" validate:"required"`
}
