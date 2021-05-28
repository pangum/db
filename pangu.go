package database

import (
	`github.com/storezhang/pangu`
)

func init() {
	app := pangu.New()

	if err := app.Provides(
		NewXormEngine,
		NewTx,
		NewRedis,
		NewElasticsearch,
	); nil != err {
		panic(err)
	}
}
