package xwf

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"go-phoenix/xwf/flow"
)

// PostStart 流程启动
func (r *Flows) PostStart(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	instanceId := ctx.PostFormValue("instanceId") // 流程实例ID
	diagramId := ctx.PostFormValue("diagramId")   // 流程ID
	values := ctx.PostFormValue("values")         // 表单数据

	// 后续节点
	var backs []ExecuteBackward
	if err := json.Unmarshal([]byte(ctx.PostFormValue("backwards")), &backs); err != nil {
		return nil, err
	}

	// 流程实例是否已启动
	var status enum.FlowStatus
	if err := asql.SelectRow(tx, "SELECT status_ FROM wf_flow WHERE instance_id_ = ?", instanceId).Scan(&status); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else if status != enum.FlowStatusRevoked {
		return nil, errors.New("流程实例已启动")
	}

	exists, executed, activated := make(map[int]struct{}), base.NewIntSet([]int{}), base.NewIntSet([]int{})
	for _, back := range backs {
		for _, route := range back.Routes {
			// 已经执行的节点忽略
			if _, ok := exists[route]; ok {
				continue
			}

			node, err := flow.NewNode(tx, ctx, diagramId, route)
			if err != nil {
				return nil, err
			}

			// Start
			if start, ok := node.(flow.StartFlowable); ok {
				executed.Append(route)

				// 是否属于撤回再次启动
				if status == enum.FlowStatusRevoked {
					if err := start.Restart(instanceId, values); err != nil {
						return nil, err
					}
				} else {
					if err := start.Start(instanceId, values); err != nil {
						return nil, err
					}
				}
			}

			// Execute Start
			if execute, ok := node.(flow.ExecuteFlowable); ok {
				if route != back.Key {
					logrus.Panic("unreachable")
				}

				activated.Append(route)
				if err := execute.ExecuteStart(instanceId, back.Executors); err != nil {
					return nil, err
				}
			}

			// Branch
			if branch, ok := node.(flow.BranchFlowable); ok {
				executed.Append(route)
				if err := branch.Branch(instanceId); err != nil {
					return nil, err
				}
			}

			// End
			if _, ok := node.(flow.EndFlowable); ok {
				logrus.Panic("unreachable")
			}

			exists[route] = struct{}{}
		}
	}

	// 更新流程实例状态
	query := "UPDATE wf_flow SET values_ = ?, executed_keys_ = ?, activated_keys_ = ?, active_at_ = ?, status_ = ? WHERE instance_id_ = ?"
	args := []interface{}{values, executed.String(), activated.String(), asql.GetNow(), enum.FlowStatusExecuting, instanceId}
	if err := asql.Update(tx, query, args...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}
