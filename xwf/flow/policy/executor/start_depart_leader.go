package flow

import (
	"database/sql"
	"go-phoenix/asql"
)

func ExecutorPolicyStartDepartLeader(tx *sql.Tx, diagramId string, key int, flowId string) ([]string, error) {
	var startDepartId string

	query := "SELECT create_depart_id_ FROM wf_flow WHERE id = ?"
	if err := asql.SelectRow(tx, query, flowId).Scan(&startDepartId); err != nil {
		return nil, err
	}

	leaders, err := GetDepartLeaderByDepartId(tx, startDepartId)
	if err != nil {
		return nil, err
	}

	return leaders, nil
}
