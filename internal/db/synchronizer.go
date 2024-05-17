package db

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/pangum/db/internal/config"
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
		field.New("models", models),
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
