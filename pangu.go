package db

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newEngine,
		newTransaction,
	)
}
