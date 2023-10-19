package db

import (
	"github.com/pangum/db/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependencies().Build().Provide(new(plugin.Creator).New)
}
