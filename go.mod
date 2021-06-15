module github.com/storezhang/pangu-database

go 1.16

require (
	github.com/storezhang/glog v1.0.7
	github.com/storezhang/gox v1.5.0
	github.com/storezhang/pangu v1.2.4
	github.com/storezhang/pangu-logging v1.0.0
	xorm.io/core v0.7.3
	xorm.io/xorm v1.1.0
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/pangu => ../pangu
