package database

import (
	`github.com/storezhang/pangu`
)

func init() {
	if err := pangu.New().Provides(newXormEngine, newTx); nil != err {
		panic(err)
	}
}
