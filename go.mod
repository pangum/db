module github.com/pangum/database

go 1.16

require (
	github.com/elliotchance/sshtunnel v1.3.1
	github.com/pangum/logging v0.0.3
	github.com/pangum/pangu v0.0.1
	github.com/storezhang/gox v1.7.9
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	xorm.io/core v0.7.3
	xorm.io/xorm v1.1.0
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/pangu => ../pangu
