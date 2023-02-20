package xwf

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"go-phoenix/xwf/flow"
	"strings"
)

type ExecuteBackward struct {
	Key       int             `json:"key"`       // 节点
	Name      string          `json:"name"`      // 名称
	Routes    []int           `json:"routes"`    // 节点路由
	Executors []flow.Executor `json:"executors"` // 执行者
}

// PostExecuteAccept 流程执行
func (o *Flows) PostExecuteAccept(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	var backs []ExecuteBackward

	id := ctx.PostFormValue("id")           // 流转节点ID
	values := ctx.PostFormValue("values")   // 表单数据
	comment := ctx.PostFormValue("comment") // 审批意见

	// 后续节点
	if err := json.Unmarshal([]byte(ctx.PostFormValue("backwards")), &backs); err != nil {
		return nil, err
	}

	// 校验数据是否合法
	var flowId, diagramId string
	var key int
	var executedKeys, activatedKeys string
	query := `
		SELECT wf_flow.flow_id_, wf_flow.diagram_id_, wf_flow_node.key_, wf_flow.executed_keys_, wf_flow.activated_keys_
		FROM wf_flow_node,wf_flow 
		WHERE wf_flow.id = wf_flow_node.flow_id_ AND wf_flow_node.id = ? 
			AND wf_flow_node.executor_user_id_ = ? AND wf_flow_node.status_ = ?
	`
	args := []interface{}{id, ctx.GetUserId(), enum.FlowNodeStatusExecuting}
	if err := asql.SelectRow(tx, query, args...).Scan(&flowId, &diagramId, &key, &executedKeys, &activatedKeys); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("没有处理该待办事项权限")
		}

		return nil, err
	}

	exists, status := make(map[int]struct{}), enum.FlowStatusExecuting
	executed, activated := base.NewIntSetFromString(executedKeys), base.NewIntSetFromString(activatedKeys)
BREAK:
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
			if _, ok := node.(flow.StartFlowable); ok {
				logrus.Panic("unreachable")
			}

			// Execute
			if execute, ok := node.(flow.ExecuteFlowable); ok {
				if route != key && route != back.Key {
					logrus.Panic("unreachable")
				}

				// End
				if route == key {
					activated.Remove(route) // 从激活节点移除
					executed.Append(route)  // 添加到已执行节点

					if err := execute.ExecuteAccept(id, values, comment); err != nil {
						return nil, err
					}
				}

				// Start
				if route == back.Key {
					activated.Append(route)
					if err := execute.ExecuteStart(flowId, back.Executors); err != nil {
						return nil, err
					}
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
			if end, ok := node.(flow.EndFlowable); ok {
				executed.Append(route)
				status = enum.FlowStatusFinished
				if err := end.End(flowId, values); err != nil {
					return nil, err
				}

				break BREAK
			}

			exists[route] = struct{}{}
		}
	}

	now := asql.GetNow()

	// 如果状态为结束，那么将激活节点作废
	if status == enum.FlowStatusFinished {
		query := "UPDATE wf_flow_node SET status_ = ?, canceled_at_ = ? WHERE instance_id_ = ? AND status_ = ?"
		args := []interface{}{enum.FlowNodeStatusCanceled, now, flowId, enum.FlowNodeStatusExecuting}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		activated.Reset()
	}

	// 流程提示信息
	statusText := "流程实例已结束"
	if status != enum.FlowStatusFinished {
		users, err := o.executors(tx, flowId)
		if err != nil {
			return nil, err
		}

		statusText = fmt.Sprintf("等待 %s 执行中", strings.Join(users, "  "))
	}

	// 更新流程状态
	queryUpdate := "UPDATE wf_flow SET values_ = ?, executed_keys_ = ?, activated_keys_ = ?, active_at_ = ?, status_ = ?, status_text_ = ? WHERE instance_id_ = ?"
	argsUpdate := []interface{}{values, executed.String(), activated.String(), now, status, statusText, flowId}
	if status == enum.FlowStatusFinished {
		queryUpdate = "UPDATE wf_flow SET values_ = ?, executed_keys_ = ?, activated_keys_ = ?, active_at_ = ?, end_at_ = ?, status_ = ?, status_text_ = ? WHERE instance_id_ = ?"
		argsUpdate = []interface{}{values, executed.String(), activated.String(), now, now, status, statusText, flowId}
	}

	if err := asql.Update(tx, queryUpdate, argsUpdate...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}
