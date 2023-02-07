package jsvm

import (
	"database/sql"
	"github.com/robertkrimen/otto"
	"go-phoenix/handle"
	vm "go-phoenix/rujs/jsvm/sql"
)

func NewVM(tx *sql.Tx, ctx *handle.Context, reg func(*otto.Otto) error) (*otto.Otto, error) {
	vmo := otto.New()

	// 请求上下文
	if err := vmo.Set("ctx", NewContext(ctx)); err != nil {
		return nil, err
	}

	// SQL相关操作
	if err := vmo.Set("sql", vm.NewSql(tx, ctx)); err != nil {
		return nil, err
	}

	// 日志输出
	if err := vmo.Set("log", NewLog()); err != nil {
		return nil, err
	}

	// 其他对象
	if reg != nil {
		if err := reg(vmo); err != nil {
			return nil, err
		}
	}

	return vmo, nil
}
