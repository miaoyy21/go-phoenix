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
	flowId := ctx.PostFormValue("flowId") // 流程实例ID

	// 后续节点
	var backs []ExecuteBackward
	if err := json.Unmarshal([]byte(ctx.PostFormValue("backwards")), &backs); err != nil {
		return nil, err
	}

	// 流程实例是否已启动
	var diagramId, values string
	var status enum.FlowStatus
	if err := asql.SelectRow(tx, "SELECT diagram_id_, values_, status_ FROM wf_flow WHERE id = ?", flowId).Scan(&diagramId, &values, &status); err != nil {
		return nil, err
	}

	// 只能启动 草稿、撤回和驳回的流程
	if status != enum.FlowStatusRevoked && status != enum.FlowStatusDraft && status != enum.FlowStatusRejected {
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
					if err := start.Start(flowId, values); err != nil {
						return nil, err
					}
				} else {
					if err := start.Start(flowId, values); err != nil {
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
				if err := execute.ExecuteStart(flowId, back.Executors); err != nil {
					return nil, err
				}
			}

			// Branch
			if branch, ok := node.(flow.BranchFlowable); ok {
				executed.Append(route)
				if err := branch.Branch(flowId); err != nil {
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
	query := "UPDATE wf_flow SET  executed_keys_ = ?, activated_keys_ = ?, active_at_ = ?, status_ = ? WHERE id = ?"
	args := []interface{}{executed.String(), activated.String(), asql.GetNow(), enum.FlowStatusExecuting, flowId}
	if err := asql.Update(tx, query, args...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}
