module github.com/pangum/database

go 1.17

require (
	github.com/elliotchance/sshtunnel v1.3.1
	github.com/pangum/logging v0.0.9
	github.com/pangum/pangu v0.0.9
	github.com/pkg/errors v0.9.1 // indirect
	github.com/goexl/gox v0.0.4
	golang.org/x/crypto v0.0.0-20211115234514-b4de73f9ece8
	gopkg.in/yaml.v2 v2.3.0 // indirect
	xorm.io/core v0.7.3
	xorm.io/xorm v1.1.0
)

// replace github.com/goexl/gox => ../gox
// replace github.com/storezhang/pangu => ../pangu
// replace github.com/pangum/logging => ../logging
