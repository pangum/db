package db

import (
	"github.com/pangum/db/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Puts(
		ctor.DB,
		ctor.New,
		ctor.NewTransaction,
		ctor.NewSynchronizer,
	).Build().Apply()
}
