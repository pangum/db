package db

import (
	"runtime"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/pangum/logging"
	"xorm.io/xorm"
)

type (
	// Tx 事务控制
	Tx struct {
		engine *Engine
		logger *logging.Logger
	}

	txFun  func(tx *Session) (err error)
	txpFun func(tx *Session, params ...interface{}) (err error)
)

// 事务控制
func newTx(engine *Engine, logger *logging.Logger) *Tx {
	return &Tx{
		engine: engine,
		logger: logger,
	}
}

func (t *Tx) Do(fun txFun, fields ...gox.Field[any]) (err error) {
	return t.do(func(tx *Session) error {
		return fun(tx)
	}, fields...)
}

func (t *Tx) Dop(fun txpFun, params []interface{}, fields ...gox.Field[any]) (err error) {
	return t.do(func(tx *Session) error {
		return fun(tx, params...)
	}, fields...)
}

func (t *Tx) do(fun func(tx *Session) error, fields ...gox.Field[any]) (err error) {
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

func (t *Tx) begin(tx *xorm.Session, fields ...gox.Field[any]) (err error) {
	if err = tx.Commit(); nil != err {
		t.error(err, "开始数据库事务出错", fields...)
	}

	return
}

func (t *Tx) commit(tx *xorm.Session, fields ...gox.Field[any]) {
	if err := tx.Commit(); nil != err {
		t.error(err, "提交数据库事务出错", fields...)
	}
}

func (t *Tx) close(tx *xorm.Session, fields ...gox.Field[any]) {
	if err := tx.Close(); nil != err {
		t.error(err, "关闭数据库事务出错", fields...)
	}
}

func (t *Tx) rollback(tx *xorm.Session, fields ...gox.Field[any]) {
	if err := tx.Rollback(); nil != err {
		t.error(err, "回退数据库事务出错", fields...)
	}
}

func (t *Tx) error(err error, msg string, fields ...gox.Field[any]) {
	fun, _, line, _ := runtime.Caller(1)

	logFields := make([]gox.Field[any], 0, len(fields)+4)
	logFields = append(logFields, field.New("fun", runtime.FuncForPC(fun).Name()))
	logFields = append(logFields, field.New("line", line))
	logFields = append(logFields, fields...)
	logFields = append(logFields, field.Error(err))
	t.logger.Error(msg, logFields...)
}
