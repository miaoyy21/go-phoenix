package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strconv"
)

type SysAutoNos struct {
}

func (o *SysAutoNos) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	code := ctx.FormValue("code")

	no, err := asql.AutoNo(tx, code, 1, ctx.Values())
	if err != nil {
		return nil, err
	}

	return no[0], nil
}

func (o *SysAutoNos) GetPatch(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	code := ctx.FormValue("code")
	sCount := ctx.FormValue("count")
	count, err := strconv.Atoi(sCount)
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("缺少指定获取数量参数【count】")
	}

	return asql.AutoNo(tx, code, count, ctx.Values())
}
