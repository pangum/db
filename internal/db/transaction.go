package db

import (
	"runtime"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"xorm.io/xorm"
)

type Transaction struct {
	engine *Engine
	logger log.Logger

	_ gox.Pointerized
}

func NewTransaction(engine *Engine, logger log.Logger) *Transaction {
	return &Transaction{
		engine: engine,
		logger: logger,
	}
}

func (t *Transaction) Do(fun Function, fields ...gox.Field[any]) (err error) {
	return t.do(func(tx *Session) error {
		return fun(tx)
	}, fields...)
}

func (t *Transaction) do(fun func(tx *Session) error, fields ...gox.Field[any]) (err error) {
	session := t.engine.NewSession()
	if err = t.begin(session, fields...); nil != err {
		return
	}
	defer t.close(session, fields...)

	if err = fun(&Session{Session: session}); nil != err {
		t.rollback(session, fields...)
	} else {
		t.commit(session, fields...)
	}

	return
}

func (t *Transaction) begin(tx *xorm.Session, fields ...gox.Field[any]) (err error) {
	if err = tx.Begin(); nil != err {
		t.error(err, "开始数据库事务出错", fields...)
	}

	return
}

func (t *Transaction) commit(tx *xorm.Session, fields ...gox.Field[any]) {
	if err := tx.Commit(); nil != err {
		t.error(err, "提交数据库事务出错", fields...)
	}
}

func (t *Transaction) close(tx *xorm.Session, fields ...gox.Field[any]) {
	if err := tx.Close(); nil != err {
		t.error(err, "关闭数据库事务出错", fields...)
	}
}

func (t *Transaction) rollback(tx *xorm.Session, fields ...gox.Field[any]) {
	if err := tx.Rollback(); nil != err {
		t.error(err, "回退数据库事务出错", fields...)
	}
}

func (t *Transaction) error(err error, msg string, fields ...gox.Field[any]) {
	fun, _, line, _ := runtime.Caller(1)

	logFields := make([]gox.Field[any], 0, len(fields)+4)
	logFields = append(logFields, field.New("fun", runtime.FuncForPC(fun).Name()))
	logFields = append(logFields, field.New("line", line))
	logFields = append(logFields, fields...)
	logFields = append(logFields, field.Error(err))
	t.logger.Error(msg, logFields...)
}
