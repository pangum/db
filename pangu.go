package db

import (
	"github.com/pangum/db/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.Config,
		ctor.Sync,
		ctor.New,
		ctor.NewTransaction,
		ctor.NewSynchronizer,
	).Build().Build().Apply()
}
