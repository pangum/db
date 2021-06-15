package database

import (
	`github.com/storezhang/pangu`
	_ `github.com/storezhang/pangu-logging`
)

func init() {
	if err := pangu.New().Provides(newXormEngine, newTx); nil != err {
		panic(err)
	}
}
