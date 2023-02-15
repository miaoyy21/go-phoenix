package flow

import (
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/rujs"
	"go-phoenix/xwf/enum"
	flow "go-phoenix/xwf/flow/policy"
	"strings"
)

func (node *NodeExecute) ScopeExecutors(flowId string, values string) ([]string, error) {
	scope := make([]string, 0, 16)

	// 如果是所有用户，那么就不进行其他判断
	if node.executorPolicy == enum.ExecutorPolicyAllUsers {
		fn := flow.ExecutorPolicyMap[enum.ExecutorPolicyAllUsers]

		return fn(node.tx, node.diagramId, node.key, flowId)
	}

	// User
	if len(node.executorUsers) > 0 {
		users := strings.Split(node.executorUsers, ",")
		scope = append(scope, users...)
	}

	// Depart
	if len(node.executorDeparts) > 0 {
		departs := strings.Split(node.executorDeparts, ",")
		scope = append(scope, departs...)
	}

	// Role
	if len(node.executorRoles) > 0 {
		roles := strings.Split(node.executorRoles, ",")

		args := make([]interface{}, 0, len(roles))
		for _, role := range roles {
			args = append(args, role)
		}

		query := "SELECT organization_id_ FROM sys_organization_role WHERE role_id_ in (?" + strings.Repeat(",?", len(roles)-1) + ")"
		res, err := asql.Select(node.tx, query, args...)
		if err != nil {
			return nil, err
		}

		rows := base.ResAsSliceString(res, "organization_id_")
		scope = append(scope, rows...)
	}

	// Executor Policy
	if node.executorPolicy != enum.ExecutorPolicyNone {
		fn, ok := flow.ExecutorPolicyMap[node.executorPolicy]
		if !ok {
			return nil, fmt.Errorf("not implement Executor Policy %s", node.executorPolicy)
		}

		// 执行者策略函数调用
		rows, err := fn(node.tx, node.diagramId, node.key, flowId)
		if err != nil {
			return nil, err
		}

		scope = append(scope, rows...)
	}

	// Executor Script
	if len(strings.TrimSpace(node.executorScript)) > 4 {
		value, err := rujs.Run(node.tx, node.ctx, node.executorScript, 0, flowReg(node, values))
		if err != nil {
			return nil, err
		}

		val, err := value.Export()
		if err != nil {
			return nil, err
		}

		executors, ok := val.([]string)
		if !ok {
			return nil, errors.New("自定义执行者脚本必须返回由组织ID组成的数组")
		}

		scope = append(scope, executors...)
	}

	return base.SliceStringAsSet(scope), nil
}
