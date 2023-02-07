package xmd

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
)

type SysDictItems struct {
}

func (m *SysDictItems) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if strings.EqualFold(scope, "ALL") {
		query = `
		SELECT sys_dict_kind.code_ AS code, sys_dict_item.code_ AS id, sys_dict_item.name_ AS value
		FROM sys_dict_kind, sys_dict_item
		WHERE sys_dict_kind.id = sys_dict_item.kind_id_
	`
	}

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
