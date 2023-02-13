package xsys

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysAutoNos struct {
}

func (o *SysAutoNos) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	code := ctx.FormValue("code")

	no, err := asql.AutoNo(tx, code, ctx.GetValues())
	if err != nil {
		return nil, err
	}

	return map[string]string{"status": "success", "no": no}, nil
}
