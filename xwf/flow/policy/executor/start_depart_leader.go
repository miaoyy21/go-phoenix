package flow

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
)

func ExecutorPolicyStartDepartLeader(tx *sql.Tx, diagramId string, key int, instanceId string) ([]string, error) {
	var startDepartId string

	query := "SELECT create_depart_id_ FROM wf_flow WHERE diagram_id_ = ? AND instance_id_ = ?"
	if err := asql.SelectRow(tx, query, diagramId, instanceId).Scan(&startDepartId); err != nil {
		return nil, err
	}

	leaders, err := GetDepartLeaderByDepartId(tx, startDepartId)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("ExecutorPolicyStartDepartLeader is %#v", leaders)

	return leaders, nil
}
