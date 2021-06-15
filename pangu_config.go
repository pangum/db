package database

type panguConfig struct {
	// 关系型数据库配置
	Database Config `json:"database" yaml:"database" validate:"required"`
}
