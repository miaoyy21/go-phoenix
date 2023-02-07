package xmd

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strings"
)

type SysSetting struct {
}

func (m *SysSetting) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	scope := ctx.FormValue("scope")

	if strings.EqualFold(scope, "LOGIN") {
		res, err := asql.Select(tx, "SELECT field_, value_ FROM sys_setting")
		if err != nil {
			return nil, err
		}

		newRes, _ := base.ResAsMapSlice(res, true, "field_", "value_")
		return newRes, nil
	}

	query := `
		SELECT id, code_, name_, description_, create_at_, update_at_
		FROM wf_diagram
		ORDER BY order_ ASC
	`
	res, err := asql.Select(tx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
