package xwf

import (
	"database/sql"
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"go-phoenix/xwf/flow"
)

// PostExecuteReject 流程驳回
func (o *Flows) PostExecuteReject(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")           // 流转节点ID
	comment := ctx.PostFormValue("comment") // 审批意见

	// 校验数据是否合法
	var flowId, diagramId string
	var key int
	var values, executedKeys, activatedKeys string
	query := `
		SELECT wf_flow.id, wf_flow.diagram_id_, wf_flow_task.key_, wf_flow.values_, wf_flow.executed_keys_, wf_flow.activated_keys_
		FROM wf_flow_task,wf_flow 
		WHERE wf_flow.id = wf_flow_task.flow_id_ AND wf_flow_task.id = ? 
			AND wf_flow_task.executor_user_id_ = ? AND wf_flow_task.status_ = ?
	`
	args := []interface{}{id, ctx.UserId(), enum.FlowNodeStatusExecuting}
	if err := asql.SelectRow(tx, query, args...).Scan(&flowId, &diagramId, &key, &values, &executedKeys, &activatedKeys); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("没有处理该待办事项权限")
		}

		return nil, err
	}

	// 节点
	node, err := flow.NewNode(tx, ctx, flowId, diagramId, key)
	if err != nil {
		return nil, err
	}

	// 必须是执行环节
	execute, ok := node.(flow.ExecuteFlowable)
	if !ok {
		return nil, errors.New("驳回只会执行环节有效")
	}

	// 是否允许驳回
	if !execute.Rejectable() {
		return nil, errors.New("不允许进行驳回操作")
	}

	// 是否填写驳回意见
	if execute.RequireRejectComment() && len(comment) < 1 {
		return nil, errors.New("请填写审批意见（驳回原因）")
	}

	// 执行驳回
	if err := execute.ExecuteReject(id, flowId, values, comment); err != nil {
		return nil, err
	}

	// 流程提示信息
	statusText := fmt.Sprintf("【%s】%s 已驳回", node.Name(), ctx.UserName())

	// 更新流程状态
	queryUpdate := "UPDATE wf_flow SET activated_keys_ = ?, active_at_ = ?, status_ = ?, status_text_ = ? WHERE id = ?"
	argsUpdate := []interface{}{base.NewIntSet([]int{key}).String(), asql.GetNow(), enum.FlowStatusRejected, statusText, flowId}
	if err := asql.Update(tx, queryUpdate, argsUpdate...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}
