package database

import (
	`github.com/pangum/pangu`
)

func init() {
	if err := pangu.New().Provides(newEngine, newTx); nil != err {
		panic(err)
	}
}
