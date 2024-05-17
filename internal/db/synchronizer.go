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
	config *config.Sync
	logger log.Logger
}

func NewSynchronizer(engine *Engine, config *config.Sync, logger log.Logger) *Synchronizer {
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
		default:
			tables = append(tables, s.getType(table))
		}
	}

	return
}

func (s *Synchronizer) getType(table any) (typ string) {
	if t := reflect.TypeOf(table); t.Kind() == reflect.Ptr {
		typ = "*" + t.Elem().Name()
	} else {
		typ = t.Name()
	}

	return
}
