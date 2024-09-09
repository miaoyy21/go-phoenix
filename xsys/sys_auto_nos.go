package xsys

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysAutoNos struct {
}

func (o *SysAutoNos) Any(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	code := ctx.FormValue("code")

	no, err := asql.AutoNo(tx, code, 1, ctx.Values())
	if err != nil {
		return nil, err
	}

	return map[string]string{"status": "success", "no": no}, nil
}
