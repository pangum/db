module github.com/storezhang/pangu-database

go 1.16

require (
	github.com/go-redis/redis/v8 v8.8.0
	github.com/olivere/elastic/v7 v7.0.24
	github.com/storezhang/glog v1.0.5
	github.com/storezhang/gox v1.4.9
	github.com/storezhang/pangu v1.1.10
	xorm.io/core v0.7.3
	xorm.io/xorm v1.0.5
)

// replace github.com/storezhang/gox => ../gox
