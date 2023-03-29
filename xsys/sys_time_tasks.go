package xsys

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysTimeTasks struct {
}

func (o *SysTimeTasks) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := "SELECT id, code_, name_, sync_status_, description_, create_at_, update_at_ FROM sys_table ORDER BY order_ ASC"

	return asql.Select(tx, query)
}
