package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
)

type SysDocs struct {
}

func (o *SysDocs) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	docs := strings.Split(ctx.FormValue("id"), ",")
	if len(docs) < 1 {
		return make([]interface{}, 0), nil
	}

	query := fmt.Sprintf("SELECT id, name_, size_ FROM sys_doc WHERE id IN (?%s)", strings.Repeat(", ?", len(docs)-1))
	args := make([]interface{}, 0, len(docs))
	for _, doc := range docs {
		args = append(args, doc)
	}

	return asql.Select(tx, query, args...)
}
