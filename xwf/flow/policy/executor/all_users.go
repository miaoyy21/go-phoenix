package flow

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/base"
)

func ExecutorPolicyAllUsers(tx *sql.Tx, diagramId string, key int, flowId string) ([]string, error) {
	query := "SELECT id FROM sys_depart ORDER BY order_ ASC"
	res, err := asql.Select(tx, query)
	if err != nil {
		return nil, err
	}

	return base.ResAsSliceString(res, "id"), nil
}
