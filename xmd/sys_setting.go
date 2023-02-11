package xmd

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
)

type SysSetting struct {
}

func (m *SysSetting) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	res, err := asql.Select(tx, "SELECT field_, value_ FROM sys_setting")
	if err != nil {
		return nil, err
	}

	newRes, _ := base.ResAsMapSlice(res, true, "field_", "value_")
	return newRes, nil
}

func (m *SysSetting) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	values := ctx.GetValues()

	for field, value := range values {
		query := "UPDATE sys_setting SET value_ = ?, update_at_ = ? WHERE field_ = ?"
		if err := asql.Update(tx, query, value, asql.GetNow(), field); err != nil {
			return nil, err
		}
	}

	return map[string]string{"status": "success"}, nil
}
