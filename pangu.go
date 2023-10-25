package db

import (
	"github.com/pangum/db/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	creator := new(plugin.Creator)
	pangu.New().Get().Dependency().Put(
		creator.New,
		creator.NewTransaction,
	).Build().Build().Apply()
}
