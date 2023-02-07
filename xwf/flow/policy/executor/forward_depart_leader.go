package flow

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/xwf/enum"
)

func ExecutorPolicyForwardDepartLeader(tx *sql.Tx, diagramId string, key int, instanceId string) ([]string, error) {
	var startDepartId string

	query := "SELECT executed_depart_id_ FROM wf_flow_node WHERE diagram_id_ = ? AND instance_id_ = ? AND (category_ = ? OR category_ = ?) ORDER BY order_ DESC"
	if err := asql.SelectRow(tx, query, diagramId, instanceId, enum.CategoryStart, enum.CategoryExecute).Scan(&startDepartId); err != nil {
		return nil, err
	}

	leaders, err := GetDepartLeaderByDepartId(tx, startDepartId)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("ExecutorPolicyForwardDepartLeader is %#v", leaders)

	return leaders, nil
}
