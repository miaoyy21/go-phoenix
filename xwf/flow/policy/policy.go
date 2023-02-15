package flow

import (
	"database/sql"
	"go-phoenix/xwf/enum"
	flow "go-phoenix/xwf/flow/policy/executor"
)

type ExecutorPolicyFunc func(tx *sql.Tx, diagramId string, key int, flowId string) ([]string, error)

var ExecutorPolicyMap = map[enum.ExecutorPolicy]ExecutorPolicyFunc{
	enum.ExecutorPolicyAllUsers:            flow.ExecutorPolicyAllUsers,
	enum.ExecutorPolicyStartDepartLeader:   flow.ExecutorPolicyStartDepartLeader,
	enum.ExecutorPolicyForwardDepartLeader: flow.ExecutorPolicyForwardDepartLeader,
}
