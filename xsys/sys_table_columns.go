package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysTableColumns struct {
}

func (o *SysTableColumns) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query, args := "invalid", make([]interface{}, 0)

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysTableColumns) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}
