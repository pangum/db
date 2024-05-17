package db

import (
	"reflect"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/pangum/db/internal/config"
	"github.com/pangum/db/internal/db/internal/core"
)

type Synchronizer struct {
	gox.CannotCopy

	engine *Engine
	config *config.DB
	logger log.Logger
}

func NewSynchronizer(engine *Engine, config *config.DB, logger log.Logger) *Synchronizer {
	return &Synchronizer{
		engine: engine,
		config: config,
		logger: logger,
	}
}

func (s *Synchronizer) Sync(models ...any) (err error) {
	fields := gox.Fields[any]{
		field.New("models", s.tables(models...)),
		field.New("config", s.config),
	}
	s.logger.Info("同步数据库表开始", fields...)
	if err = s.engine.Sync(models...); nil == err {
		s.logger.Info("同步数据库表成功", fields...)
	} else {
		s.logger.Error("同步数据库表失败", fields...)
	}

	return
}

func (s *Synchronizer) tables(models ...any) (tables []string) {
	tables = make([]string, 0, len(models))
	for _, model := range models {
		switch table := model.(type) {
		case core.Commenter:
			tables = append(tables, table.TableComment())
		case core.NameMaker:
			tables = append(tables, table.TableName())
		default:
			tables = append(tables, s.config.TableMapper().Obj2Table(s.getType(table)))
		}
	}

	return
}

func (s *Synchronizer) getType(table any) (typ string) {
	if of := reflect.TypeOf(table); of.Kind() == reflect.Ptr {
		typ = of.Elem().Name()
	} else {
		typ = of.Name()
	}

	return
}
